package routes

import (
	"payment-system/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine) {
	transactionRoutes := r.Group("/transactions")
	{
		transactionRoutes.Use(controllers.AuthMiddleware()) // Secure transactions with JWT
		transactionRoutes.POST("/create", controllers.CreateTransaction)
		transactionRoutes.GET("/:id", controllers.GetTransactionByID) // ✅ Fetch transactions
	}
}
