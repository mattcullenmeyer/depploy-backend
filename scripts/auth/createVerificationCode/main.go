package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	otpPtr := flag.String("c", "", "verification code")
	usernamePtr := flag.String("u", "", "username")
	emailPtr := flag.String("e", "", "email")
	flag.Parse()

	createUserArgs := authModel.CreateVerificationCodeParams{
		Otp:      *otpPtr,
		Username: *usernamePtr,
		Email:    *emailPtr,
	}

	if err := authModel.CreateVerificationCode(createUserArgs); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully created verification code")
}
