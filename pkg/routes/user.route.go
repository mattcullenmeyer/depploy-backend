package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/pkg/controllers"
)

func UserRoute(router *gin.Engine) {
	user := router.Group("/user")

	user.GET("/:username", controllers.Username)
}
