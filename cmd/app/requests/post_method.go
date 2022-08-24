package requests

import (
	"net/http"

	"errors"

	"strconv"

	"math/rand"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"github.com/gin-contrib/sessions"

	"go-web-app-experiment/cmd/app/notebook_db"

	"go-web-app-experiment/cmd/app/user_session"

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
			db.Delete(&notebook_db.Note{}, id)

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

		// Get User ID Data
		user_id := user_session.GetUserSessionData(c, "user_id").(uint)

		// Create new note entry
		notebook_db.CreateNewNoteEntry(db, title, description, user_id)

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
		notebook_db.UpdateNoteEntry(db, id, title, description)

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
func RegisterNewUser(db *gorm.DB, email string, username string, password string, confirm string) error {
	/* Check if email is already taken */
	var user1 = &notebook_db.User{}

	db.Where("email = ?", email).First(&user1)

	if email == user1.Email {
		return errors.New("Email already taken")
	}

	/* Check if username is already taken */
	var user2 = &notebook_db.User{}

	db.Where("username = ?", username).First(&user2)

	if username == user2.Username {
		return errors.New("Username already taken")
	}

	/* Check if the password is under 6 characters */
	if len(password) < 6 {
		return errors.New("Password is less than 6 characters")
	} 

	/* Checks if the passwords match */
	if password != confirm {
		return errors.New("Passwords do not match")
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

// Not Post

func IsUserValid(db *gorm.DB, username string, password string) (uint, error) {
	var user = &notebook_db.User{}

	// Get entry with the specified email or username
	db.Where("email = ? OR username = ?", username, username).First(&user)

	if username == user.Email || username == user.Username {
		// Check if the email or username exists 
		if password != user.Password {
			// Check if password is incorrect
			return 0, errors.New("Password is incorrect")
		}
	} else if username != user.Email || username != user.Username {
		// Check if the email or username does not exists 
		return 0, errors.New("Username does not exist")
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
		user_id, err := IsUserValid(db, username, password)

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


