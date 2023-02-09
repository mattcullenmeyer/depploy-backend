package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuperAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		superAdmin := c.GetBool("superAdmin")

		if !superAdmin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this resource"})
			return
		}

		c.Next()
	}
}
