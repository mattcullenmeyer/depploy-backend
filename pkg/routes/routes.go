package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/pkg/middleware"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	middleware.CorsMiddleware(router)

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	AuthRoute(router.Group("/auth"))
	UserRoute(router.Group("/user"))

	return router
}
