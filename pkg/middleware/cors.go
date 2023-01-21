package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

// https://github.com/gin-contrib/cors
func CorsMiddleware(router *gin.Engine) {
	allowOriginsString, err := utils.GetParameter("CorsAllowOrigins") // TODO: Error handling and consolidate getting parameters
	if err != nil {
		log.Println(err)
	}

	allowOrigins := strings.Split(allowOriginsString, ",")

	config := cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))
}
