package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		// Authorization header format is "Bearer <token>"
		values := strings.Split(header, " ")

		if len(values) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "You are not logged in"})
			return
		}

		name, token := values[0], values[1]

		if name != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "You are not logged in"})
			return
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		if !claims.Authorized {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not authorized to access this resource"})
			return
		}

		// TODO: Check to make sure user exists in the database

		c.Set("username", claims.Username)
		c.Set("account", claims.Account)

		c.Next()
	}
}
