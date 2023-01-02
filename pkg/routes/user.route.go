package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/pkg/controllers"
	"github.com/mattcullenmeyer/depploy-backend/pkg/middleware"
)

func UserRoute(router *gin.RouterGroup) {
	router.GET("/:username", controllers.Username)
	router.GET("/user", middleware.TokenAuth(), controllers.GetUser)
}
