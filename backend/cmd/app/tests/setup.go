package tests

import (
	"github.com/gin-gonic/gin"
	"notebook_app/cmd/app/notebook_db"
	"notebook_app/cmd/app/router_setup"
)

func Setup() *gin.Engine {
	/* Setup the app for unit testing */
	// Open database
	db := notebook_db.UseSQLite("database/notebook.db")

	// setup router
	r := router_setup.InitializeRouter(db)

	return r
}

