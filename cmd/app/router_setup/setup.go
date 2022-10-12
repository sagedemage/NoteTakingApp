package router_setup

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notes"

	"notebook_app/cmd/app/auth"
)

func InitializeRouter(db *gorm.DB) *gin.Engine {
	// create the router
	router := gin.Default()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	// html renderer
	//router.HTMLRender = template_loader.LoadTemplates("cmd/app/templates")

	// session
	//store := cookie.NewStore([]byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))

	// do not trust all proxies for security reasons
	router.SetTrustedProxies(nil)

	// Load static files (for css, and etc)
	router.Static("/static", "cmd/app/static")

	/* API */
	api := router.Group("/api")

	// check user authentication
	//api.GET("/check-user-auth", auth.CheckUserAuthenticated)

	// Register the user
	api.POST("/register", auth.Register(db))

	// login user
	api.POST("/login", auth.Login(db))

	// logout user
	api.GET("/logout", auth.Logout)

	// fetch user's notes
	api.POST("/view-notes", notes.ViewNotes(db))

	// add a note
	api.POST("/add-new-note", notes.AddNewNote(db))

	// Delete Note from POST request
	api.POST("/delete-note", notes.DeleteNote(db))

	// Fetch Note
	api.POST("/fetch-note", notes.FetchNote(db))

	// Edit Note from POST request
	api.POST("/edit-note", notes.EditNote(db))

	/* Old routes for my purely backend app */
	/*
	// Render the home page at the root of the website
	router.GET("/", page_renderer.RenderWebPage("home.tmpl", "Home"))

	// Render the about page at the route "/about"
	router.GET("/about", page_renderer.RenderWebPage("about.tmpl", "About"))

	// Render the new registration page at route "/register"
	router.GET("/register", auth.RegisterPage)

	// Render the login page at route "/login"
	router.GET("/login", auth.LoginPage)

	// Register the user
	router.POST("/register", auth.Register123(db))

	// Login the user
	router.POST("/login", auth.Login123(db))

	// Page Not Found
	router.NoRoute(page_renderer.RenderPageNotFoundWebPage("404.html", "404 Page - Page Not Found"))

	// Auhtorization Required
	auth_routes := router.Group("/").Use(auth.AuthRequired)

	// Render the view table page at route "/table"
	auth_routes.GET("/dashboard", notes.ViewNotes123(db))

	// Render the new entry page at route "/add-new-note"
	auth_routes.GET("/add-new-note", page_renderer.RenderWebPage("add-note.tmpl", "New Note"))

	// Render the new entry page at route "/edit-note"
	auth_routes.GET("/edit-note", notes.EditNoteForm123(db))

	// Render the new entry page at route "/delete-entry"
	auth_routes.GET("/delete-note", page_renderer.RenderWebPage("delete-note.tmpl", "Delete Note"))

	// Logout the user
	auth_routes.GET("/logout", auth.Logout123)

	// Handle Delete and Edit post requets
	auth_routes.POST("/dashboard", notes.DeleteOrEditNote(db))

	// Add Note from POST request
	auth_routes.POST("/add-new-note", notes.AddNewNote123(db))

	// Edit Note from POST request
	auth_routes.POST("/edit-note", notes.EditNote123(db))

	// Delete Note from POST request
	auth_routes.POST("/delete-note", notes.DeleteNote123(db))
	*/

	return router
}
