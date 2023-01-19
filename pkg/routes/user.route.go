package routes

import (
	"github.com/gin-gonic/gin"
	userController "github.com/mattcullenmeyer/depploy-backend/pkg/controllers/user"
	"github.com/mattcullenmeyer/depploy-backend/pkg/middleware"
)

func UserRoute(router *gin.RouterGroup) {
	// router.GET("/username/:username", userController.Username) // TODO: Clean up Username controller
	router.GET("/email/:email", userController.CheckEmailAvailability)
	router.GET("/user", middleware.TokenAuth(), userController.GetUser)
	router.GET("/users", userController.GetUsers)
}
