package main

import (
	"github.com/gin-gonic/gin"

	"go-web-app-experiment/cmd/app/models"

	"go-web-app-experiment/cmd/app/requests"
)

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

	/* Get Requests */
	// Render the home page at the root of the website
	router.GET("/", requests.RenderWebPage("home/index.tmpl", "Home"))

	// Render the about page at the route "/about"
	router.GET("/about", requests.RenderWebPage("about/index.tmpl", "About"))

	// Render the new registration page at route "/register"
	router.GET("/register", requests.RenderWebPage("auth/register.tmpl", "Register"))

	// Render the login page at route "/login"
	router.GET("/login", requests.RenderWebPage("auth/login.tmpl", "Login"))

	// Render the view table page at route "/table"
	router.GET("/view-notes", requests.ViewNotes(db))

	// Render the new entry page at route "/new-entry"
	router.GET("/add-new-note", requests.RenderWebPage("notes/new-note.tmpl", "New Note"))

	// Render the new entry page at route "/new-entry"
	router.GET("/edit-note", requests.EditNoteForm(db))

	/* Post Requests */
	// Render the view table page at route "/table"
	router.POST("/view-notes", requests.DeleteOrEditNote(db)) //

	// Get Form data from POST request
	router.POST("/add-new-note", requests.AddNewNote(db)) //

	// Get Form data from POST request
	router.POST("/edit-note", requests.EditNote(db)) //

	// listen and serve on localhost:8080
	router.Run(":8080")
}
