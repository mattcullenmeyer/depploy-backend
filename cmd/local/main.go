package main

import (
	"log"

	"github.com/mattcullenmeyer/depploy-backend/router"
)

func runRouter() {
	router := router.RegisterRoutes()

	err := router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runRouter()
}
