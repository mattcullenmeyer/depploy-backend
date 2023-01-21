package userController

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func GetUsers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "25")
	lastEvaluatedKey := c.Query("next")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit query parameter must be a number"})
		return
	}

	result, err := userModel.FetchUsers(limit, lastEvaluatedKey)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": result.Users, "next": result.Next})
}
