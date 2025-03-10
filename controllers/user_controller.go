package controllers

import (
	"net/http"
	"payment-system/config"
	"payment-system/models"
	"payment-system/services"

	"github.com/gin-gonic/gin"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := services.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registerd successfully", "user": user})

}

// Login USER authenticate user and returns a JWT token
func LoginUser(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := services.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := config.GenerateToken((user.ID))
	c.JSON(http.StatusOK, gin.H{"token": token})
}
