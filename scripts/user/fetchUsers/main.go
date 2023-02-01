package main

import (
	"flag"
	"fmt"
	"strconv"

	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func main() {
	limitPtr := flag.String("l", "10", "limit")
	lastKeyPtr := flag.String("k", "", "last evaluated key")
	flag.Parse()

	limit, _ := strconv.ParseInt(*limitPtr, 10, 64)

	fetchUsersArgs := userModel.FetchUsersParams{
		Limit: limit,
		Key:   *lastKeyPtr,
	}

	result, err := userModel.FetchUsers(fetchUsersArgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", result)
}
