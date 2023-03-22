package routes

import (
	"github.com/gin-gonic/gin"
	userController "github.com/mattcullenmeyer/depploy-backend/pkg/controllers/user"
	"github.com/mattcullenmeyer/depploy-backend/pkg/middleware"
)

// WARNING: Do NOT add endpoint to change SuperAdmin access for security reasons
func UserRoute(router *gin.RouterGroup) {
	// router.GET("/username/:username", userController.Username) // TODO: Clean up Username controller
	router.GET("/email/:email", userController.CheckEmailAvailability)
	router.GET("/details", middleware.TokenAuth(), userController.GetUser)
	router.GET("/users", middleware.TokenAuth(), middleware.SuperAdminAuth(), userController.GetUsers)
	router.PATCH("/password", middleware.TokenAuth(), userController.UpdatePassword)
	router.PATCH("/email", middleware.TokenAuth(), userController.UpdateEmail)
}
