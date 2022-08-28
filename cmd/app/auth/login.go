package auth

import (
	"net/http"
	
	"gorm.io/gorm"
	
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"

	"golang.org/x/crypto/bcrypt"

	"go-web-app-experiment/cmd/app/notebook_db"

	"go-web-app-experiment/cmd/app/form"
)

func is_user_valid(db *gorm.DB, username string, password string) (uint, error) {
	/* Check if the User is Valid */
	var user = &notebook_db.User{}

	// Get entry with the specified email or username
	db.Where("email = ? OR username = ?", username, username).First(&user)

	if username == user.Email || username == user.Username {
		// Check if the email or username exists 
		// compare the password to the password hash
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil {
			// Check if the password is incorrect
			return 0, errors.New("password is incorrect") 
		}
	} else if username != user.Email || username != user.Username {
		// Check if the email or username does not exists 
		return 0, errors.New("username does not exist")
	}

	return user.ID, nil
}

func Login(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Login */
		// Parse Form Data
		c.Request.ParseForm()

		// get username form data
		var username string = form.GetFormValue(c, "username") 
		
		// get password form data
		var password string = form.GetFormValue(c, "password") 

		// Is User Valid
		user_id, err := is_user_valid(db, username, password)

		/* Check if user registration is successful */
		if err == nil {
			token := GenerateSessionToken()

			c.SetCookie("token", token, 3600, "", c.Request.URL.Hostname(), false, true)
			
			// user session
			session := sessions.Default(c)

			// store that logged in is true
			session.Set("is_logged_in", true)
			session.Set("user_id", user_id)
      		session.Save()

			// Redirect to the table view page
			c.Redirect(http.StatusFound, "/view-notes")
		} else {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "Login Failed",
				"ErrorMessage": err.Error(),
			})
		}

		
	}
	return gin.HandlerFunc(fn)
}

func Logout(c *gin.Context) {
	/* Logout */
	// user session
	session := sessions.Default(c)

	// delete the stored logged in variable
	session.Delete("is_logged_in")
	session.Delete("user_id")
	session.Save()

	// Redirect to the table view page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

