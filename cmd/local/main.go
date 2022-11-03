package main

import (
	"github.com/mattcullenmeyer/depploy-backend/router"
)

func runRouter() {
	router := router.RegisterRoutes()

	router.Run(":8080")
}

func main() {
	runRouter()
}
