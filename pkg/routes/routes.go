package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/pkg/controllers"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.GET("/secret", controllers.GetSecret)

	AuthRoute(router.Group("/auth"))
	UserRoute(router.Group("/user"))

	return router
}
