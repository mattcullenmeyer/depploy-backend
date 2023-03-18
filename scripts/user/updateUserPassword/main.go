package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	accountIdPtr := flag.String("a", "", "account ID")
	passwordPtr := flag.String("p", "", "password")
	flag.Parse()

	updateUserPasswordArgs := userModel.UpdateUserPasswordParams{
		AccountId: *accountIdPtr,
		Password:  *passwordPtr,
	}

	if err := userModel.UpdateUserPassword(updateUserPasswordArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated account password")
}
