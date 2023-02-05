package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	accountPtr := flag.String("a", "", "account id")
	accessPtr := flag.Bool("s", false, "superuser admin access boolean")
	flag.Parse()

	updateUserSuperuserArgs := userModel.UpdateUserSuperuserParams{
		AccountId: *accountPtr,
		Access:    *accessPtr,
	}

	if err := userModel.UpdateUserSuperuser(updateUserSuperuserArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated superuser admin access")
}
