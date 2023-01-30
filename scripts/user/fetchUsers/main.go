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
	lastEvaluatedKey := *lastKeyPtr

	result, err := userModel.FetchUsers(limit, lastEvaluatedKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", result)
}
