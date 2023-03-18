package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	otpPtr := flag.String("p", "", "password")
	flag.Parse()

	deleteOtpArgs := authModel.DeleteOtpParams{
		Otp: *otpPtr,
	}

	result, err := authModel.DeleteOtp(deleteOtpArgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", result)
}
