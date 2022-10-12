package auth

import (

	"gorm.io/gorm"

	"errors"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/data_types"

)

func is_user_valid(db *gorm.DB, username string, password string) (uint, error) {
	/* Check if the User is Valid */
	var user = &notebook_db.User{}

	// Get entry with the specified email or username
	db.Where("email = ? OR username = ?", username, username).First(&user)

	// Incorrect username or password (Reddit, GitHub)

	if username == user.Email || username == user.Username {
		// Check if the email or username exists
		// compare the password to the password hash
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil {
			// Check if the password is incorrect
			return user.ID, errors.New("incorrect username or password")
		}
	} else if username != user.Email || username != user.Username {
		// Check if the email or username does not exists
		return user.ID, errors.New("incorrect username or password")
	}

	return user.ID, nil
}



func Login(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Login */

		type RequestBody struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var body RequestBody

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
			// send login error message
			//login_error_message(c, err)

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

/* Old Functions for the Purely Backend App */
// for legacy reasons

/*
func LoginPage(c *gin.Context) {
	/ Login Page /
	var page_title = "Login"

	var error_status, success_status = false, false

	// query msg_error
	msg_error := c.Query("msg_error")

	// query msg_success
	msg_success := c.Query("msg_success")

	// Get user session data
	session := sessions.Default(c)

	// Get loggin in value
	user := session.Get("is_logged_in")

	var http_status = http.StatusOK

	if msg_error != "" {
		http_status = http.StatusUnprocessableEntity
		error_status = true
	} else if msg_success != "" {
		success_status = true
	}
	c.HTML(http_status, "login.tmpl", gin.H{
		"page_title":     page_title,
		"user":           user,
		"error_status":   error_status,
		"msg_error":      msg_error,
		"success_status": success_status,
		"msg_success":    msg_success,
	})
}

func Login123(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/ Login /
		// Parse Form Data
		c.Request.ParseForm()

		// get username form data
		var username string = form.GetFormValue(c, "username")

		// get password form data
		var password string = form.GetFormValue(c, "password")

		// Is User Valid
		user_id, err := is_user_valid(db, username, password)

		/ Check if user registration is successful /
		if err == nil {
			// user session
			session := sessions.Default(c)

			// store that logged in is true
			session.Set("is_logged_in", true)
			session.Set("user_id", user_id)
			session.Save()

			// Redirect to the dashboard
			c.Redirect(http.StatusFound, "/dashboard")
		} else {
			// send login error message
			login_error_message(c, err)
		}
	}
	return gin.HandlerFunc(fn)
}

func Logout123(c *gin.Context) {
	/ Logout /
	// user session
	session := sessions.Default(c)

	// delete the stored logged in variable
	session.Delete("is_logged_in")
	session.Delete("user_id")
	session.Save()

	// Redirect to the table view page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func login_error_message(c *gin.Context, err error) {
	// initialize query values
	q := url.Values{}

	// set note_id query value
	q.Set("msg_error", err.Error())

	// pass query value to the delete note route
	location := url.URL{Path: "/login", RawQuery: q.Encode()}

	// redirect to edit note
	c.Redirect(http.StatusFound, location.RequestURI())
}
*/
