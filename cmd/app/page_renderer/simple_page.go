package page_renderer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

func RenderWebPage(template_page string, page_title string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Get user session data
		session := sessions.Default(c)

		// Get loggin in value
		user := session.Get("is_logged_in")

		c.HTML(http.StatusOK, template_page, gin.H{
			"page_title": page_title,
			"user": user,
		})
	}
	return gin.HandlerFunc(fn)
}

