package auth

import (
	"github.com/gin-gonic/gin"

	//"notebook_app/cmd/app/user_session"
)


//c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

func CheckUserAuthenticated(c *gin.Context) {
	// Get user logged_in session data
	//user := user_session.GetUserSessionData(c, "is_logged_in")

	//token, err := c.Cookie("token")


	//token := c.Request.Header["token"]

	token := c.GetHeader("token")

	/*if err != nil {
		println(err)
	}*/
	
	// would be a boolean value [true or false]
	c.JSON(200, gin.H{
		"token": token,
	})
}
