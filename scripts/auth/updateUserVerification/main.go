package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	usernamePtr := flag.String("u", "", "username")
	verifiedPtr := flag.Bool("v", true, "verified boolean")
	flag.Parse()

	username := *usernamePtr
	verified := *verifiedPtr

	if err := authModel.UpdateUserVerified(username, verified); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated user verification")
}
