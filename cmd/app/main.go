package main

import (
	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/router_setup"
)

func main() {
	// Open database
	db := notebook_db.InitDB("database/notebook.db")

	// Initialize the router
	r := router_setup.InitializeRouter(db)

	// listen and serve on localhost:8080
	r.Run(":8080")
}
