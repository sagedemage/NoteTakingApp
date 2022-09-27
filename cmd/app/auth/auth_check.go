package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

func CheckUserAuthenticated(c *gin.Context) {
	// Get user session data
	session := sessions.Default(c)

	// Get login in value
	user := session.Get("is_logged_in")
	
	// would be a boolean value [true or false]
	c.JSON(200, gin.H{
		"is_logged_in": user,
	})
}
