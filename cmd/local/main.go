package main

import (
	"log"

	"github.com/mattcullenmeyer/depploy-backend/pkg/routes"
)

func runRouter() {
	router := routes.RegisterRoutes()

	err := router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runRouter()
}
