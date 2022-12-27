package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattcullenmeyer/depploy-backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
// See "Model binding and validation" section
func RegisterUser(c *gin.Context) {
	var registration models.Registration

	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	hash, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}

	// Check to make sure email / username doesn't already exist
	// Generate verification code (and encode) and then send verification email

	c.JSON(http.StatusOK, gin.H{"password": registration.Password, "hash": hash})
}
