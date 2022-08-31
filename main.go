package main

func main() {
	router := registerRoutes()

	router.Run(":8080")
}
