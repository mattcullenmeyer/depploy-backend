package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	otpPtr := flag.String("p", "", "one-time password")
	accountIdPtr := flag.String("a", "", "account id")
	emailPtr := flag.String("e", "", "email")
	flag.Parse()

	createOtpArgs := authModel.CreateOtpParams{
		Otp:       *otpPtr,
		AccountId: *accountIdPtr,
		Email:     *emailPtr,
	}

	if err := authModel.CreateOtp(createOtpArgs); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully created one-time password")
}
