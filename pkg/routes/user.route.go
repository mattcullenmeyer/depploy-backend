package routes

import (
	"github.com/gin-gonic/gin"
	userController "github.com/mattcullenmeyer/depploy-backend/pkg/controllers/user"
	"github.com/mattcullenmeyer/depploy-backend/pkg/middleware"
)

// WARNING: Do NOT add endpoint to change superuser access for security reasons
func UserRoute(router *gin.RouterGroup) {
	// router.GET("/username/:username", userController.Username) // TODO: Clean up Username controller
	router.GET("/username/:username", userController.CheckUsernameAvailability)
	router.GET("/details", middleware.TokenAuth(), userController.GetUser)
	router.GET("/users", middleware.TokenAuth(), middleware.SuperuserAuth(), userController.GetUsers)
	router.PATCH("/username", middleware.TokenAuth(), userController.UpdateUsername)
}
