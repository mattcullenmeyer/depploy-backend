package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	otpPtr := flag.String("c", "", "verification code")
	accountPtr := flag.String("a", "", "account id")
	usernamePtr := flag.String("u", "", "username")
	emailPtr := flag.String("e", "", "email")
	flag.Parse()

	createUserArgs := authModel.CreateVerificationCodeParams{
		Otp:       *otpPtr,
		AccountId: *accountPtr,
		Username:  *usernamePtr,
		Email:     *emailPtr,
	}

	if err := authModel.CreateVerificationCode(createUserArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully created verification code")
}
