package main

import (
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/users"
)

func registerRoutes() *ginadapter.GinLambda {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/:username", users.Profile)
	}

	return ginadapter.New(router)
}
