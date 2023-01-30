package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	usernamePtr := flag.String("u", "", "username")
	emailPtr := flag.String("e", "", "email")
	passwordPtr := flag.String("p", "", "password")
	flag.Parse()

	createUserArgs := authModel.CreateEmailUserParams{
		Username: *usernamePtr,
		Email:    *emailPtr,
		Password: *passwordPtr,
	}

	if err := authModel.CreateEmailUser(createUserArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully created new user")
}
