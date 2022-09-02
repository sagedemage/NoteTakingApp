package page_renderer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	c.AbortWithStatus(http.StatusUnauthorized)
}
