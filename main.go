package main

import (
	"github.com/mattcullenmeyer/depploy-backend/routes"
)

func main() {
	router := routes.RegisterRoutes()

	router.Run(":8080")
}