package main

import (
	"github.com/mattcullenmeyer/depploy-backend/router"
)

func runRouter() {
	router := router.RegisterRoutes()

	err := router.Run(":8080")

	if err != nil {
		// do something here
	}
}

func main() {
	runRouter()
}
