package auth

import (
	"github.com/gin-gonic/gin"

	"notebook_app/cmd/app/user_session"

	"notebook_app/cmd/app/page_renderer"
)


// Middleware to check the user session
func AuthRequired(c *gin.Context) {


	// Get user logged_in session data
	user := user_session.GetUserSessionData(c, "is_logged_in")

	if user == nil {
		page_renderer.RenderUnauthorizedWebPage(c)

		return
	} 
}

