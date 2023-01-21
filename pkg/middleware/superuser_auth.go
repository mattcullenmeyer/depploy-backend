package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuperuserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		superuser := c.GetBool("superuser")

		if !superuser {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this resource"})
			return
		}

		c.Next()
	}
}
