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

	// do not trust all proxies for security reasons
	router.SetTrustedProxies(nil)

	/* API */
	api := router.Group("/api")

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

	// Get Decoded Token from POST request
	api.POST("/get-decoded-token", auth.GetDecodedToken)

	return router
}
