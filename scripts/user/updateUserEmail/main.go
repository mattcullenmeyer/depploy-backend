package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	accountIdPtr := flag.String("a", "", "account ID")
	emailPtr := flag.String("e", "", "email")
	flag.Parse()

	updateUserEmailArgs := userModel.UpdateUserEmailParams{
		AccountId: *accountIdPtr,
		Email:     *emailPtr,
	}

	if err := userModel.UpdateUserEmail(updateUserEmailArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated account email")
}
