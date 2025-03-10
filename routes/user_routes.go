package routes

import (
	"payment-system/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.POST("/login", controllers.LoginUser)
	}
}
