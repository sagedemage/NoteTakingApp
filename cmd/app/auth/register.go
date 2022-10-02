package auth

import (
	"net/http"

	"net/url"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"errors"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/form"
	
	"notebook_app/cmd/app/data_types"
)

func RegisterPage(c *gin.Context) {
	/* Login Page */
	var page_title = "Register"

	// success message status
	var error_status = false

	// query msg_success
	msg_error := c.Query("msg_error")

	if msg_error != "" {
		error_status = true
	}

	// Get user session data
	session := sessions.Default(c)

	// Get loggin in value
	user := session.Get("is_logged_in")

	c.HTML(http.StatusOK, "register.tmpl", gin.H{
		"page_title": page_title,
		"user": user,
		"error_status": error_status,
		"msg_error": msg_error,
	})
}

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
		err := register_new_user(db, email, username, password, confirm)

		/* Check if user registration is successful */
		if err == nil {
			// Send success message
			register_success_message(c, "Registered Successfully")
		} else {
			// Send error message
			register_error_message(c, err)
		}
	}
	return gin.HandlerFunc(fn)
}

func Register2(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Register */
		// Parse Form Data
		/*c.Request.ParseForm()

		// get email data
		var email string = form.GetFormValue(c, "email") 
		
		// get username form data
		var username string = form.GetFormValue(c, "username") 
		
		// get password form data
		var password string = form.GetFormValue(c, "password") 

		// get confirm from data
		var confirm string = form.GetFormValue(c, "confirm")*/

		type RequestBody struct {
			Email string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
			ConfirmPwd string `json:"confirm_pwd"`
		}

		var body RequestBody

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
			//register_success_message(c, "Registered Successfully")
			c.JSON(200, data_types.JSON{
				"registered":  true,
				"msg_success": "Registered Successfully",
			})
		} else {
			// Send error message
			register_error_message(c, err)

			// json message
			c.JSON(200, data_types.JSON{ // 401
				"registered": false, 
				"msg_error": err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}

func register_error_message(c *gin.Context, err error) {
	// initialize query values
	q := url.Values{}

	// set note_id query value
	q.Set("msg_error", err.Error())

	// pass query value to the delete note route
	location := url.URL{Path: "/register", RawQuery: q.Encode()}

	// redirect back to register page
	c.Redirect(http.StatusFound, location.RequestURI())
}

func register_success_message(c *gin.Context, msg_success string) {
	// initialize query values
	q := url.Values{}

	// set note_id query value
	q.Set("msg_success", msg_success)

	// pass query value to the delete note route
	location := url.URL{Path: "/login", RawQuery: q.Encode()}

	// redirect to edit note
	c.Redirect(http.StatusFound, location.RequestURI())
}
