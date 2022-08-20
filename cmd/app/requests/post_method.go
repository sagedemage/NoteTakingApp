package requests

import (
	"net/http"

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

		/* Check if email is already taken */
		var user1 = &models.User{}

		db.Where("email = ?", email).First(&user1)

		if email == user1.Email {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "email already taken",
				"Email": user1.Email,
			})
			return
		}

		/* Check if username is already taken */
		var user2 = &models.User{}

		db.Where("username = ?", username).First(&user2)

		if username == user2.Username {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "username already taken",
				"Username": user2.Username,
			})
			return
		}

		/* Check if the password is over 6 characters */
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "password must be greater than 6 characters",
			})
			return
		} 

		/* Checks if the passwords is match */
		if password != confirm {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "passwords do not match",
			})
			return
		}

		// Create user account
		db.Create(&models.User{Email: email, Username: username, Password: password})

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/login")
	}
	return gin.HandlerFunc(fn)
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

		var user = &models.User{}

		// Get entry with the specified email or username
		db.Where("email = ? OR username = ?", username, username).First(&user)

		if username == user.Email || username == user.Username {
			// Check if the email or username exists 
			if password != user.Password {
				// Check if password is incorrect
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"MSG": "Password is incorrect",
				})
			return
			}
		} else if username != user.Email || username != user.Username {
			// Check if the email or username does not exists 
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"MSG": "username does not exist",
				"Email": username,
			})
			return
		} 

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")
	}
	return gin.HandlerFunc(fn)
}
