package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web-app-experiment/cmd/app/models"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.LoadHTMLGlob("cmd/app/templates/**/*")

	router.Static("/static", "cmd/app/static/")

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

	// Open database
	db := models.OpenDatabase("database/test.db")

	// Migrate the schema
	db.AutoMigrate(&models.Product{})

	// Create 
	db.Create(&models.Product{Name: "Shampoo", Price: 20})

	// Read
	var product models.Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "Shampoo") // find product with name Shampoo

	// Update - upgrade product's price to 200
	db.Model(&product).Update("Price", 200)

	// Update - update multiple fields
	db.Model(&product).Updates(models.Product{Price: 200, Name: "Shampoo"}) // non-zero fields

	// Delete - delete product
	db.Delete(&product, 1)

	// listen and serve on localhost:8080
	router.Run(":8080") 
}
