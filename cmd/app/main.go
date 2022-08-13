package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web-app-experiment/cmd/app/models"
)

func main() {
	// create the router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("cmd/app/templates/**/*")

	// Load static files (for css, and etc)
	router.Static("/static", "cmd/app/static/")

	// Open database
	db := models.OpenDatabase("database/test.db")

	// Migrate the schema
	db.AutoMigrate(&models.Product{})

	// Render the home page at the root of the website
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tmpl", gin.H {
			"title": "Home Page",
		})
	})

	// Render the about page at the route "/about"
	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about/index.tmpl", gin.H {
			"title": "About Page",
		})
	})

	// Create 
	db.Create(&models.Product{Name: "Shampoo", Price: 20})

	// Product model
	var product models.Product
	
	/* Read single object from the database */
	// find the product with integer primary key
	db.First(&product, 1) 
	// find product with name Shampoo
	db.First(&product, "code = ?", "Shampoo") 

	// Update - upgrade product's price to 200
	db.Model(&product).Update("Price", 200)

	// Update - update multiple fields
	db.Model(&product).Updates(models.Product{Price: 200, Name: "Shampoo"}) // non-zero fields

	// Delete - delete product
	db.Delete(&product, 1)

	// listen and serve on localhost:8080
	router.Run(":8080") 
}
