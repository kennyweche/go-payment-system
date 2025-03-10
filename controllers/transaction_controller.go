package controllers

import (
	"net/http"
	"payment-system/config"
	"payment-system/models"
	"payment-system/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		claims, err := config.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user_id", (*claims)["user_id"])
		c.Next()
	}
}

// CreateTransaction publishes a transaction to Kafka
func CreateTransaction(c *gin.Context) {
	var req models.Transaction
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Publish transaction to Kafka
	if err := services.PublishTransaction(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to queue transaction"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Transaction queued for processing"})
}

// GetTransactionByID retrieves a transaction
func GetTransactionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := services.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}
