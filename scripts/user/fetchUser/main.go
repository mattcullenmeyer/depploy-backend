package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	usernamePtr := flag.String("u", "", "username")
	flag.Parse()

	username := *usernamePtr

	result, err := userModel.FetchUser(username)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", result)
}
