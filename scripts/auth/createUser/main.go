package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	accountIdPtr := flag.String("a", "", "account id")
	emailPtr := flag.String("e", "", "email")
	flag.Parse()

	createUserArgs := authModel.CreateEmailUserParams{
		AccountId: *accountIdPtr,
		Email:     *emailPtr,
	}

	if err := authModel.CreateEmailUser(createUserArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully created new user")
}
