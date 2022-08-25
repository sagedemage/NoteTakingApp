package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web-app-experiment/cmd/app/user_session"
)


// Middleware to check the user session
func AuthRequired(c *gin.Context) {
	// Get user logged_in session data
	user := user_session.GetUserSessionData(c, "is_logged_in")

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unathorized"})
		return
	} 
}

