package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	accountPtr := flag.String("a", "", "account id")
	accessPtr := flag.Bool("s", false, "super admin access boolean")
	flag.Parse()

	updateUserSuperAdminArgs := userModel.UpdateUserSuperAdminParams{
		AccountId: *accountPtr,
		Access:    *accessPtr,
	}

	if err := userModel.UpdateUserSuperAdmin(updateUserSuperAdminArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated super admin access")
}
