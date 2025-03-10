package utils

import (
	"github.com/gin-gonic/gin"
)

// JSONResponse formats API responses
func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

// ErrorResponse simplifies error handling
func ErrorResponse(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}
