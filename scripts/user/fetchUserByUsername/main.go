package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	usernamePtr := flag.String("u", "", "username")
	flag.Parse()

	fetchUserByUsernameArgs := userModel.FetchUserByUsernameParams{
		Username: *usernamePtr,
	}

	result, err := userModel.FetchUserByUsername(fetchUserByUsernameArgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", result)
}
