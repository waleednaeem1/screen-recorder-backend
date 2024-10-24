package routes

import (
	"screen-recorder-backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/forgot-password", controllers.ForgotPassword) // New forgot password route
	}
}
