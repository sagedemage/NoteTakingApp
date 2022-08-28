package auth

import (
	//"net/http"

	"github.com/gin-gonic/gin"

	"go-web-app-experiment/cmd/app/user_session"

	"go-web-app-experiment/cmd/app/page_renderer"
)


// Middleware to check the user session
func AuthRequired(c *gin.Context) {
	// Get user logged_in session data
	user := user_session.GetUserSessionData(c, "is_logged_in")

	if user == nil {
		page_renderer.RenderUnauthorizedWebPage(c, "401.html", "401 Page - Unauthorized")
		//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unathorized"})
		//return
	} 
}

