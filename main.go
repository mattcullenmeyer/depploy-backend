package main

func main() {
	router := RegisterRoutes()

	router.Run(":8080")
}