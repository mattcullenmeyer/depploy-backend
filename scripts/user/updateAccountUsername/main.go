package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	usernamePtr := flag.String("u", "", "username")
	accountIdPtr := flag.String("a", "", "account ID")
	flag.Parse()

	updateAccountUsernameArgs := userModel.UpdateAccountUsernameParams{
		Username:  *usernamePtr,
		AccountId: *accountIdPtr,
	}

	if err := userModel.UpdateAccountUsername(updateAccountUsernameArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated account username")
}
