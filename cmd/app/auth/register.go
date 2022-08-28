package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"gorm.io/gorm"
	
	"errors"
	
	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/form"
)

func RegisterNewUser(db *gorm.DB, email string, username string, password string, confirm string) error {
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
	if len(password) < 6 {
		return errors.New("password is less than 6 characters")
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
		// Parse Form Data
		c.Request.ParseForm()

		// get email data
		var email string = form.GetFormValue(c, "email") 
		
		// get username form data
		var username string = form.GetFormValue(c, "username") 
		
		// get password form data
		var password string = form.GetFormValue(c, "password") 

		// get confirm from data
		var confirm string = form.GetFormValue(c, "confirm") 

		// Register new user
		err := RegisterNewUser(db, email, username, password, confirm)

		/* Check if user registration is successful */
		if err == nil {
			token := GenerateSessionToken()

			c.SetCookie("token", token, 3600, "", c.Request.URL.Hostname(), false, true)
			c.Set("is_logged_in", true)

			// Redirect to the login page
			c.Redirect(http.StatusFound, "/login")
		} else {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "Registration Failed",
				"ErrorMessage": err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}

