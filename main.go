package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")

	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tmpl", gin.H {
			"title": "Home Page",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about/index.tmpl", gin.H {
			"title": "About Page",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
