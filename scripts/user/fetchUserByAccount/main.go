package main

import (
	"flag"
	"fmt"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	accountPtr := flag.String("a", "", "account id")
	flag.Parse()

	fetchUserByAccountArgs := userModel.FetchUserByAccountParams{
		AccountId: *accountPtr,
	}

	result, err := userModel.FetchUserByAccount(fetchUserByAccountArgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", result)
}
