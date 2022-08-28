package main

import (
	"go-web-app-experiment/cmd/app/router_setup"
)

func main() {
	// Initialize the router
	r := router_setup.InitializeRouter()

	// listen and serve on localhost:8080
	r.Run(":8080")
}
