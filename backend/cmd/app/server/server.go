package server

import (
	"notebook_app/cmd/app/notebook_db"
	"notebook_app/cmd/app/router_setup"
)

func Start() {
	// Open database
	db := notebook_db.InitDB()
	
	// Initialize the router
	router := router_setup.InitializeRouter(db)
	
	// listen and serve on localhost:8080
	router.Run(":8080")
}
