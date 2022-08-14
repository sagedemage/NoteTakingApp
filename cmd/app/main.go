package main

import (
	"net/http"
	"strconv"

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
	db.AutoMigrate(&models.Note{})

	// Product model
	//var note models.Note

	/* Read single object from the database */
	// find the product with integer primary key
	//db.First(&note, 1) 
	// find product with name Shampoo
	//db.First(&note, "code = ?", "Shampoo") 

	// Update - upgrade product's price to 200
	//db.Model(&note).Update("Best Shampoo", "johnson")

	// Update - update multiple fields
	//db.Model(&note).Updates(models.Note{Title: "Best Shampoos", Description: "Johnson Dove"}) // non-zero fields

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
	router.GET("/view-notes", func(c *gin.Context) {
		/* View all the database entries as a table */
		// entries of the product database
		notes := models.GetDatabase(db)

		c.HTML(http.StatusOK, "notes/view-notes.tmpl", gin.H {
			"title": "Table",
			"list": notes,
		})
	})

	// Render the view table page at route "/table"
	router.POST("/view-notes", func(c *gin.Context) {
		/* View all the database entries as a table */

		// Parse Form Data
		c.Request.ParseForm()

		var id int
		var err error

		if c.Request.PostFormValue("delete") != "" {
			/* Get Title and secription from the Post request */
			id, err = strconv.Atoi(c.Request.PostFormValue("delete"))
			if err != nil {
				panic(err)
			}
			// delete entry
			db.Delete(&models.Note{}, id)

			// redirect to notes view page
			c.Redirect(http.StatusFound, "/view-notes")
		} else if c.Request.PostFormValue("edit") != "" {
			id, err = strconv.Atoi(c.Request.PostFormValue("edit"))
			if err != nil {
				panic(err)
			}
			// delete entry
			//db.Delete(&models.Note{}, id)

			// redirect to edit note
			c.Redirect(http.StatusFound, "/edit-note")
		}
	})

	// Render the new entry page at route "/new-entry"
	router.GET("/add-new-note", func(c *gin.Context) {
		/* View all the database entries as a table */
		// entries of the product database

		c.HTML(http.StatusOK, "notes/new-note.tmpl", gin.H {
			"title": "New Note",
		})
	})

	// Get Form data from POST request
	router.POST("/add-new-note", func(c *gin.Context) {
		// Parse Form Data
		c.Request.ParseForm()
		
		/* Get Title and secription from the Post request */
		var title string = c.Request.PostFormValue("title")
		var description string = c.Request.PostFormValue("description")
	
		// Create entry // the issue
		db.Create(&models.Note{Title: title, Description: description})

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")
	})

	// Render the new entry page at route "/new-entry"
	router.GET("/edit-note", func(c *gin.Context) {
		/* View all the database entries as a table */
		// entries of the product database

		c.HTML(http.StatusOK, "notes/edit-note.tmpl", gin.H {
			"title": "Edit Note",
		})
	})

	// Get Form data from POST request
	router.POST("/edit-note", func(c *gin.Context) {
		// Parse Form Data
		c.Request.ParseForm()
		
		/* Get Title and secription from the Post request */
		var title string = c.Request.PostFormValue("title")
		var description string = c.Request.PostFormValue("description")
	
		// Create entry // the issue
		db.Create(&models.Note{Title: title, Description: description})

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")
	})

	// listen and serve on localhost:8080
	router.Run(":8080") 
}
