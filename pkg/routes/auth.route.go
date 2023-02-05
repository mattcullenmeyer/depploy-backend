package routes

import (
	"github.com/gin-gonic/gin"
	authController "github.com/mattcullenmeyer/depploy-backend/pkg/controllers/auth"
)

func AuthRoute(router *gin.RouterGroup) {
	router.POST("/register", authController.RegisterEmailUser)
	router.POST("/login", authController.LoginUser)
	router.POST("/token/refresh", authController.RefreshToken)
	router.PATCH("/email/verify", authController.VerifyEmail)
	router.POST("/email/resend", authController.ResendEmail)
	router.GET("/google", authController.GoogleOAuth)
}
