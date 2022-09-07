package page_renderer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"notebook_app/cmd/app/user_session"
)

func RenderWebPage(template_page string, page_title string) gin.HandlerFunc {
	/* Render a simple web page */
	fn := func(c *gin.Context) {
		// Get user logged_in session data
		user := user_session.GetUserSessionData(c, "is_logged_in")

		c.HTML(http.StatusOK, template_page, gin.H{
			"page_title": page_title,
			"user": user,
		})
	}
	return gin.HandlerFunc(fn)
}

