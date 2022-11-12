package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"errors"
	"notebook_app/cmd/app/notebook_db"
	"notebook_app/cmd/app/data_types"
	"notebook_app/cmd/app/request_bodies"
)

func register_new_user(db *gorm.DB, email string, username string, password string, confirm string) error {
	/* Check if email is already taken */
	var user1 = &notebook_db.User{}

	db.Where("email = ?", email).First(&user1)

	if email == user1.Email {
		return errors.New("email already taken")
	}

	/* Check if username is already taken */
	var user2 = &notebook_db.User{}

	db.Where("username = ?", username).First(&user2)

	if username == user2.Username {
		return errors.New("username already taken")
	}

	/* Check if the password is under 6 characters */
	if len(password) < 8 {
		return errors.New("must be at least 8 characters")
	} 

	/* Checks if the passwords match */
	if password != confirm {
		return errors.New("passwords do not match")
	}
		
	// Create user account
	notebook_db.CreateNewUser(db, email, username, password)
		
	return nil
}

func Register(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Register */
		var body request_bodies.RegisterRequest

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// Register new user
		err = register_new_user(db, body.Email, body.Username, body.Password, body.ConfirmPwd)

		/* Check if user registration is successful */
		if err == nil {
			// Send success message
			c.JSON(200, data_types.JSON{
				"registered":  true,
				"msg_success": "Registered Successfully",
			})
		} else {
			// Send error message
			c.JSON(200, data_types.JSON{ // 401
				"registered": false, 
				"msg_error": err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}

