package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	emailPtr := flag.String("e", "", "email")
	flag.Parse()

	fetchUserByEmailArgs := userModel.FetchUserByEmailParams{
		Email: *emailPtr,
	}

	result, err := userModel.FetchUserByEmail(fetchUserByEmailArgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", result)
}
