package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	accountPtr := flag.String("a", "", "account id")
	verifiedPtr := flag.Bool("v", true, "verified boolean")
	flag.Parse()

	updateUserVerifiedArgs := authModel.UpdateUserVerifiedParams{
		AccountId: *accountPtr,
		Verified:  *verifiedPtr,
	}

	if err := authModel.UpdateUserVerified(updateUserVerifiedArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated user verification")
}
