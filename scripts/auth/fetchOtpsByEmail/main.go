package main

import (
	"flag"
	"fmt"

	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
)

func main() {
	emailPtr := flag.String("e", "", "email")
	flag.Parse()

	fetchOtpsByEmailArgs := authModel.FetchOtpsByEmailParams{
		Email: *emailPtr,
	}

	result, err := authModel.FetchOtpsByEmail(fetchOtpsByEmailArgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", result)
}
