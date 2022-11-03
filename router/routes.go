package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/users"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "new deployment"})
	})

	user := router.Group("/user")
	{
		user.GET("/:username", users.Profile)
	}

	return router
}
