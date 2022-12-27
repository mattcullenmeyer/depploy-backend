package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/pkg/controllers"
)

func AuthRoute(router *gin.Engine) {
	auth := router.Group("/auth")

	auth.POST("/register", controllers.RegisterUser)
}
