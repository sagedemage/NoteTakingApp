package auth

import (

	"gorm.io/gorm"

	"errors"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/data_types"
	
	"notebook_app/cmd/app/request_bodies"

)

func is_user_valid(db *gorm.DB, username string, password string) (uint, error) {
	/* Check if the User is Valid */
	var user = &notebook_db.User{}

	// Get entry with the specified email or username
	db.Where("email = ? OR username = ?", username, username).First(&user)

	if username == user.Email || username == user.Username {
		/* Check if the email or username exists */
		// compare the password to the password hash
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil {
			/* Check if the password is incorrect */
			return user.ID, errors.New("incorrect username or password")
		}
	} else if username != user.Email || username != user.Username {
		/* Check if the email or username does not exists */
		return user.ID, errors.New("incorrect username or password")
	}

	return user.ID, nil
}



func Login(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Login */
		var body request_bodies.LoginRequest

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// Is User Valid
		user_id, err := is_user_valid(db, body.Username, body.Password)

		/* Check if user registration is successful */
		if err == nil {
			//serialized := user.Serialize()
			//token, _ := generateToken(serialized)

			c.JSON(200, data_types.JSON{
				"user_id":  user_id,
				"auth":  true,
			})
		} else {
			// json message
			c.JSON(200, gin.H{"auth": false, "msg_error": err.Error()})
		}
	}
	return gin.HandlerFunc(fn)
}

func Logout(c *gin.Context) {
	/* Logout */

	c.JSON(200, data_types.JSON{
		"auth": false,
	})
}

