package main

import (
	"payment-system/config"
	"payment-system/consumers"
	"payment-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize all configurations
	config.InitDB()
	config.InitRedis()
	config.InitKafka()

	// Start Kafka consumer in a separate goroutine
	go consumers.StartTransactionConsumer()

	// Start the Gin router
	r := gin.Default()

	// Setup routes
	routes.UserRoutes(r)
	routes.TransactionRoutes(r)

	// Start the server
	r.Run(":8080")
}
