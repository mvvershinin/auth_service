package routes

import (
	"auth_service/api/handlers"
	"auth_service/api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	server := gin.Default()

	authGroup := server.Group("/auth")
	{
		authGroup.POST("/login", handlers.Login)
		authGroup.GET("/verify", middlewares.AuthMiddleware(), handlers.Verify)
		// Add more routes as needed
	}

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
