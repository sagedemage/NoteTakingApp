package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web-app-experiment/cmd/app/models"
)

func main() {
	// create the router
	router := gin.Default()
	
	// do not trust all proxies for security reasons
	router.SetTrustedProxies(nil)

	// Load HTML templates
	router.LoadHTMLGlob("cmd/app/templates/**/*")

	// Load static files (for css, and etc)
	router.Static("/static", "cmd/app/static/")

	// Open database
	db := models.OpenDatabase("database/test.db")

	// Migrate the schema
	db.AutoMigrate(&models.Product{})

	// Product model
	var product models.Product

	// Create 
	//db.Create(&models.Product{Name: "Shampoo", Price: 30})
	
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
	//db.Delete(&product, 1)

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

	// Render the view table page at route "/table"
	router.GET("/view-table", func(c *gin.Context) {
		/* View all the database entries as a table */
		// entries of the product database
		products := models.GetDatabase(db)

		c.HTML(http.StatusOK, "table/view.tmpl", gin.H {
			"title": "Table",
			"list": products,
		})
	})

	// Render the new entry page at route "/new-entry"
	router.GET("/add-new-entry", func(c *gin.Context) {
		/* View all the database entries as a table */
		// entries of the product database

		c.HTML(http.StatusOK, "table/new-note.tmpl", gin.H {
			"title": "New Note",
		})
	})

	router.POST("/add-note", func(c *gin.Context) {
		title := "title"
		description := "description"

		c.JSON(200, gin.H{
			"title": title,
			"description": description,
		})
		
	})

	// listen and serve on localhost:8080
	router.Run(":8080") 
}
