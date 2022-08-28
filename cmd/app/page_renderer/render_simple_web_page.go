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

func RenderPageNotFoundWebPage(template_page string, page_title string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Render 404 web page */
		c.HTML(http.StatusNotFound, template_page, gin.H{
			"page_title": page_title,
		})
	}
	return gin.HandlerFunc(fn)
}

func RenderUnauthorizedWebPage(c *gin.Context, template_page string, page_title string) {
	/* Render 401 web page */
	c.HTML(http.StatusUnauthorized, template_page, gin.H{
		"page_title": page_title,
	})
}
