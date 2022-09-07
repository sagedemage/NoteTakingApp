package page_renderer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderPageNotFoundWebPage(c *gin.Context) {
	/* Render 404 web page */
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"page_title": "404 Page - Page Not Found",
	})
}

func RenderUnauthorizedWebPage(c *gin.Context) {
	/* Render 401 web page */
	c.HTML(http.StatusUnauthorized, "401.html", gin.H{
		"page_title": "Unauthorized",
	})

	c.AbortWithStatus(http.StatusUnauthorized)
}
