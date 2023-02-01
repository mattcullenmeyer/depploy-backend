package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	otpPtr := flag.String("c", "", "verification code")
	flag.Parse()

	fetchVerificationCodeArgs := authModel.FetchVerificationCodeParams{
		Otp: *otpPtr,
	}

	result, err := authModel.FetchVerificationCode(fetchVerificationCodeArgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", result)
}
