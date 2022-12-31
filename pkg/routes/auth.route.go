package routes

import (
	"github.com/gin-gonic/gin"
	authController "github.com/mattcullenmeyer/depploy-backend/pkg/controllers/auth"
)

func AuthRoute(router *gin.RouterGroup) {
	router.POST("/register", authController.RegisterUser)
	router.POST("/login", authController.LoginUser)
	router.POST("/verify", authController.VerifyEmail)
}
