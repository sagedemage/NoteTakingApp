package requests

import (
	"net/http"

	"errors"

	"strconv"

	"math/rand"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"go-web-app-experiment/cmd/app/models"

	"go-web-app-experiment/cmd/app/form"
)

/* Post Requests */

func DeleteOrEditNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Handle Delete and Edit Button POST Requests
		on the view notes page */
		// Parse Form Data
		c.Request.ParseForm()

		if form.GetFormValue(c, "delete") != "" {
			/* Delete Note Post request */
			// get entry id for the deleting an entry
			id := form.GetFormValue(c, "delete")

			// delete entry
			db.Delete(&models.Note{}, id)

			// redirect to notes view page
			c.Redirect(http.StatusFound, "/view-notes")
		} else if form.GetFormValue(c, "edit") != "" {
			/* Edit Note Post Request Redirect */
			// get entry id for the editing an entry
			id := form.GetFormValue(c, "edit")

			// set id of entry to edit
			c.SetCookie("id", id, 10, "/edit-note", c.Request.URL.Hostname(), false, true)

			// redirect to edit note
			c.Redirect(http.StatusFound, "/edit-note")
		}
	}
	return gin.HandlerFunc(fn)
}

func AddNewNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Parse Form Data
		c.Request.ParseForm()

		/* Get Title and description from the Post request */
		var title string = form.GetFormValue(c, "title")
		var description string = form.GetFormValue(c, "description")

		// Create entry
		db.Create(&models.Note{Title: title, Description: description})

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")
	}
	return gin.HandlerFunc(fn)
}

func EditNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Edit Note Post Request */
		// Parse Form Data
		c.Request.ParseForm()

		/* Get Title and secription from the Post request */
		var title string = form.GetFormValue(c, "title")
		var description string = form.GetFormValue(c, "description")

		// Get cookie of the id value
		id, err := c.Cookie("id")
		if err != nil {
			panic(err)
		}

		// Update the entry title and description by id
		models.UpdateNoteEntry(db, id, title, description)

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")

	}
	return gin.HandlerFunc(fn)
}


// Not post
func GenerateSessionToken() string {
	// Warning: do not use in production
	return strconv.FormatInt(rand.Int63(), 16)
}

// Not post
func RegisterNewUser(db *gorm.DB, email string, username string, password string, confirm string) (*gorm.DB, error) {
	/* Check if email is already taken */
		var user1 = &models.User{}

		db.Where("email = ?", email).First(&user1)

		if email == user1.Email {
			return nil, errors.New("Email already taken")
		}

		/* Check if username is already taken */
		var user2 = &models.User{}

		db.Where("username = ?", username).First(&user2)

		if username == user2.Username {
			return nil, errors.New("Username already taken")
		}

		/* Check if the password is under 6 characters */
		if len(password) < 6 {
			return nil, errors.New("Password is less than 6 characters")
		} 

		/* Checks if the passwords match */
		if password != confirm {
			return nil, errors.New("Passwords do not match")
		}

		// Create user account
		db.Create(&models.User{Email: email, Username: username, Password: password})
		
		return db, nil
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
		_, err := RegisterNewUser(db, email, username, password, confirm)

		//var sameSiteCookie http.SameSite;

		/* Check if user registration is successful */
		if err == nil {
			token := GenerateSessionToken()

			c.SetCookie("token", token, 3600, "", c.Request.URL.Hostname(), false, true)
			c.Set("is_logged_in", true)

			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "Successful Registration",
			})
		} else {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "Registration Failed",
				"ErrorMessage": err.Error(),
			})
		}

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/login")
	}
	return gin.HandlerFunc(fn)
}

func IsUserValid(db *gorm.DB, username string, password string) error {
	var user = &models.User{}

	// Get entry with the specified email or username
	db.Where("email = ? OR username = ?", username, username).First(&user)

	if username == user.Email || username == user.Username {
		// Check if the email or username exists 
		if password != user.Password {
			// Check if password is incorrect
			return errors.New("Password is incorrect")
		}
	} else if username != user.Email || username != user.Username {
		// Check if the email or username does not exists 
		return errors.New("Username does not exist")
	}
	return nil
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
		err := IsUserValid(db, username, password)

		/* Check if user registration is successful */
		if err == nil {
			token := GenerateSessionToken()

			c.SetCookie("token", token, 3600, "", c.Request.URL.Hostname(), false, true)
			c.Set("is_logged_in", true)

			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "Successful Login",
			})
		} else {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "Login Failed",
				"ErrorMessage": err.Error(),
			})
		}

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")
	}
	return gin.HandlerFunc(fn)
}
