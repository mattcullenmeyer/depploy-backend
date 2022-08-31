package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/users"
)

func registerRoutes() *gin.Engine {

	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/:username", users.Profile)
	}

	return router
}
