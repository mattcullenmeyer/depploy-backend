package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/pkg/controllers"
)

func UserRoute(router *gin.RouterGroup) {
	router.GET("/:username", controllers.Username)
}
