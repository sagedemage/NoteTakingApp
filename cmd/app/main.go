package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"go-web-app-experiment/cmd/app/models"

	"go-web-app-experiment/cmd/app/form"
)

/* Get Requests */

func render_web_page(template_page string, page_title string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		//var template_page string = "home/index.tmpl"
		//var page_title string = "Home Page"

		c.HTML(http.StatusOK, template_page, gin.H{
			"page_title": page_title,
		})
	}
	return gin.HandlerFunc(fn)
}

func view_notes(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* View all the database entries as a table */
		// entries of the notes database
		notes := models.GetNoteEntries(db)

		// Pass the list of notes to the web page
		c.HTML(http.StatusOK, "notes/index.tmpl", gin.H{
			"page_title": "Notes",
			"note_list":  notes,
		})
	}
	return gin.HandlerFunc(fn)
}

func edit_note_form(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Render edit note form */

		// Get cookie of the id value
		id, err := c.Cookie("id")
		if err != nil {
			panic(err)
		}

		// get entry note values
		note := models.GetNoteEntry(db, id)

		// pass the note's title and description to the form page
		c.HTML(http.StatusOK, "notes/edit-note.tmpl", gin.H{
			"page_title":       "Edit Note",
			"note_title":       note.Title,
			"note_description": note.Description,
		})

	}
	return gin.HandlerFunc(fn)
}

/* Post Requests */

func delete_or_edit_note(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
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
			/* Edit Note Post Request Redirect */
			// get entry id for the editing an entry
			id := form.GetFormValue(c, "edit")

			// set id of entry to edit
			c.SetCookie("id", id, 10, "/edit-note", c.Request.URL.Hostname(), false, true)

			// redirect to edit note
			c.Redirect(http.StatusFound, "/edit-note")
		}
	}
	return gin.HandlerFunc(fn)
}

func add_new_note(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Parse Form Data
		c.Request.ParseForm()

		/* Get Title and description from the Post request */
		var title string = form.GetFormValue(c, "title")
		var description string = form.GetFormValue(c, "description")

		// Create entry // the issue
		db.Create(&models.Note{Title: title, Description: description})

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")
	}
	return gin.HandlerFunc(fn)
}

func edit_note(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
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
		models.UpdateNoteEntry(db, id, title, description)

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")

	}
	return gin.HandlerFunc(fn)
}



func main() {
	// create the router
	router := gin.Default()

	// do not trust all proxies for security reasons
	router.SetTrustedProxies(nil)

	// Load HTML templates
	router.LoadHTMLGlob("cmd/app/templates/**/*")

	// Load static files (for css, and etc)
	router.Static("/static", "cmd/app/static")

	// Open database
	db := models.InitDB("database/test.db")

	// Render the home page at the root of the website
	router.GET("/", render_web_page("home/index.tmpl", "Home"))

	// Render the about page at the route "/about"
	router.GET("/about", render_web_page("about/index.tmpl", "About"))

	// Render the new registration page at route "/register"
	router.GET("/register", render_web_page("auth/register.tmpl", "Register"))

	// Render the login page at route "/login"
	router.GET("/login", render_web_page("auth/login.tmpl", "Login"))

	// Render the view table page at route "/table"
	router.GET("/view-notes", view_notes(db))

	// Render the view table page at route "/table"
	router.POST("/view-notes", delete_or_edit_note(db))

	// Render the new entry page at route "/new-entry"
	router.GET("/add-new-note", render_web_page("notes/new-note.tmpl", "New Note"))

	// Get Form data from POST request
	router.POST("/add-new-note", add_new_note(db))

	// Render the new entry page at route "/new-entry"
	router.GET("/edit-note", edit_note_form(db))

	// Get Form data from POST request
	router.POST("/edit-note", edit_note(db))

	// listen and serve on localhost:8080
	router.Run(":8080")
}
