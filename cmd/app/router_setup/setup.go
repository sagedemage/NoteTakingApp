package router_setup

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notes"

	"notebook_app/cmd/app/template_loader"

	"notebook_app/cmd/app/page_renderer"

	"notebook_app/cmd/app/auth"
)

func InitializeRouter(db *gorm.DB) *gin.Engine {
	// create the router
	router := gin.Default()

	// html renderer
	router.HTMLRender = template_loader.LoadTemplates("cmd/app/templates")

	// session
	store := cookie.NewStore([]byte("secret"))
  	router.Use(sessions.Sessions("mysession", store))

	// do not trust all proxies for security reasons
	router.SetTrustedProxies(nil)

	// Load static files (for css, and etc)
	router.Static("/static", "cmd/app/static")

	/* Get Requests */
	// Render the home page at the root of the website
	router.GET("/", page_renderer.RenderWebPage("home.tmpl", "Home")) // home.tmpl does not exist

	// Render the about page at the route "/about"
	router.GET("/about", page_renderer.RenderWebPage("about.tmpl", "About"))

	// Render the new registration page at route "/register"
	router.GET("/register", auth.RegisterPage)

	// Render the login page at route "/login"
	router.GET("/login", auth.LoginPage)

	/* Post Requests */
	// Register the user
	router.POST("/register", auth.Register(db))

	// Login the user
	router.POST("/login", auth.Login(db))

	// Page Not Found
	router.NoRoute(page_renderer.RenderPageNotFoundWebPage("404.html", "404 Page - Page Not Found"))

	/* Auhtorization Required */
	auth_routes := router.Group("/").Use(auth.AuthRequired)

	/* Get Requets */
	// Render the view table page at route "/table"
	auth_routes.GET("/dashboard", notes.ViewNotes(db))

	// Render the new entry page at route "/add-new-note"
	auth_routes.GET("/add-new-note", page_renderer.RenderWebPage("add-note.tmpl", "New Note"))

	// Render the new entry page at route "/edit-note"
	auth_routes.GET("/edit-note", notes.EditNoteForm(db))

	// Render the new entry page at route "/delete-entry"
	auth_routes.GET("/delete-note", page_renderer.RenderWebPage("delete-note.tmpl", "Delete Note"))

	// Logout the user
	auth_routes.GET("/logout", auth.Logout)

	/* Post Requests */
	// Handle Delete and Edit post requets
	auth_routes.POST("/dashboard", notes.DeleteOrEditNote(db))

	// Add Note from POST request
	auth_routes.POST("/add-new-note", notes.AddNewNote(db))

	// Edit Note from POST request
	auth_routes.POST("/edit-note", notes.EditNote(db))

	// Delete Note from POST request
	auth_routes.POST("/delete-note", notes.DeleteNote(db))

	return router
}
