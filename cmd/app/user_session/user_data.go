package user_session

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

func GetUserSessionData(c *gin.Context, name string) interface{} {
	/* Get session data of cookie name */
	// Get user session data
	session := sessions.Default(c)

	// Get session cookie name value
	user := session.Get(name)

	return user
}
