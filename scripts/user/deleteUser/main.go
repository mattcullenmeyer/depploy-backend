package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	accountPtr := flag.String("a", "", "account id")
	flag.Parse()

	deleteUserArgs := userModel.DeleteUserParams{
		AccountId: *accountPtr,
	}

	err := userModel.DeleteUser(deleteUserArgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully deleted user account")
}
