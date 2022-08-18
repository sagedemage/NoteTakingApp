package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web-app-experiment/cmd/app/models"

	"go-web-app-experiment/cmd/app/form"
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
		// entries of the notes database
		notes := models.GetDatabase(db)

		// Pass the list of notes to the web page
		c.HTML(http.StatusOK, "notes/index.tmpl", gin.H {
			"title": "Notes",
			"list": notes,
		})
	})

	// Render the view table page at route "/table"
	router.POST("/view-notes", func(c *gin.Context) {
		/* Handle Delete and Edit Button POST Requests 
		on the view notes page */

		// Parse Form Data
		c.Request.ParseForm()

		if form.GetFormValue(c, "delete") != "" {
			/* Delete Note Post request */

			// get entry id for the deleting an entry
			id := form.GetFormValue(c, "delete")

			// delete entry
			db.Delete(&models.Note{}, id)

			// redirect to notes view page
			c.Redirect(http.StatusFound, "/view-notes")
		} else if form.GetFormValue(c, "edit") != "" {
			/* Edit Note Post Request */

			// get entry id for the editing an entry
			id := form.GetFormValue(c, "edit")

			// set id of entry to edit
			c.SetCookie("id", id, 10, "/edit-note", c.Request.URL.Hostname(), false, true)

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
		
		/* Get Title and description from the Post request */
		var title string = form.GetFormValue(c, "title")
		var description string = form.GetFormValue(c, "description")
	
		// Create entry // the issue
		db.Create(&models.Note{Title: title, Description: description})

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")
	})

	// Render the new entry page at route "/new-entry"
	router.GET("/edit-note", func(c *gin.Context) {
		/* Render edit note form */

		// Get cookie of the id value
		id, err := c.Cookie("id")
		if err != nil {
			panic(err)
		}

		// get entry note values
		var note = models.GetEntry(db, id)

		// pass the note's title and description to the form page
		c.HTML(http.StatusOK, "notes/edit-note.tmpl", gin.H {
			"page_title": "Edit Note",
			"note_title": note.Title,
			"note_description": note.Description,
		})
	})

	// Get Form data from POST request
	router.POST("/edit-note", func(c *gin.Context) {
		/* Edit Note Post Request */
		// Parse Form Data
		c.Request.ParseForm()
		
		/* Get Title and secription from the Post request */
		var title string = form.GetFormValue(c, "title")
		var description string = form.GetFormValue(c, "description")

		// Get cookie of the id value
		id, err := c.Cookie("id")
		if err != nil {
			panic(err)
		}

		// Update the entry title and description by id
		models.UpdateEntry(db, id, title, description)

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")
	})

	// listen and serve on localhost:8080
	router.Run(":8080") 
}
