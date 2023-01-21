package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	usernamePtr := flag.String("u", "", "username")
	accessPtr := flag.Bool("a", false, "superuser admin access boolean")
	flag.Parse()

	username := *usernamePtr
	access := *accessPtr

	if err := userModel.UpdateUserSuperuser(username, access); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully updated superuser admin access")
}
