package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"

	"go-web-app-experiment/cmd/app/models"

	"go-web-app-experiment/cmd/app/requests"
)

// Middleware to check the user session
func AuthRequired(c *gin.Context) {
	// Get user session data
	session := sessions.Default(c)

	// Get loggin in value
	user := session.Get("is_logged_in")

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unathorized"})
		return
	} 
}

func main() {
	// create the router
	router := gin.Default()

	// session
	store := cookie.NewStore([]byte("secret"))
  	router.Use(sessions.Sessions("mysession", store))

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

	auth_routes := router.Group("/").Use(AuthRequired)

	// Render the view table page at route "/table"
	auth_routes.GET("/view-notes", requests.ViewNotes(db))

	// Render the new entry page at route "/new-entry"
	auth_routes.GET("/add-new-note", requests.RenderWebPage("notes/new-note.tmpl", "New Note"))

	// Render the new entry page at route "/new-entry"
	auth_routes.GET("/edit-note", requests.EditNoteForm(db))

	/* Post Requests */
	// Register the user
	router.POST("/register", requests.Register(db))

	// Login the user
	router.POST("/login", requests.Login(db))

	// Render the view table page at route "/table"
	auth_routes.POST("/view-notes", requests.DeleteOrEditNote(db))

	// Get Form data from POST request
	auth_routes.POST("/add-new-note", requests.AddNewNote(db))

	// Get Form data from POST request
	auth_routes.POST("/edit-note", requests.EditNote(db))

	// listen and serve on localhost:8080
	router.Run(":8080")
}
