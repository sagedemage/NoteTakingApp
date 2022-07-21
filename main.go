package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-experiment/models"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")

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

	// Open database
	db := models.OpenDatabase("test.db")

	// Migrate the schema
	db.AutoMigrate(&models.Product{})

	// Create 
	db.Create(&models.Product{Code: "D42", Price: 100})

	// Read
	var product models.Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - upgrade product's price to 200
	db.Model(&product).Update("Price", 200)

	// Update - update multiple fields
	db.Model(&product).Updates(models.Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
