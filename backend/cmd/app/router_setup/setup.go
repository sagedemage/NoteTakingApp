package router_setup

import (
	"notebook_app/cmd/app/auth"
	"notebook_app/cmd/app/notes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// fetch user's notes
	api.POST("/view-notes", notes.ViewNotes(db))

	// add a note
	api.POST("/add-new-note", notes.AddNewNote(db))

	// Delete Note from POST request
	api.DELETE("/delete-note", notes.DeleteNote(db))

	// Fetch Note
	// url = /api/fetch-note?id={number}
	api.GET("/fetch-note", notes.FetchNote(db))

	// Edit Note from POST request
	api.PATCH("/edit-note", notes.EditNote(db))

	// Get Decoded Token from POST request
	api.POST("/get-decoded-token", auth.GetDecodedToken)

	return router
}
