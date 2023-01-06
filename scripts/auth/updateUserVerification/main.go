package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	usernamePtr := flag.String("u", "", "username")
	flag.Parse()

	username := *usernamePtr

	if err := authModel.UpdateUserVerified(username); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully updated user verification")
}
