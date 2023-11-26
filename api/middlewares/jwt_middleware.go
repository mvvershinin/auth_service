package middlewares

import (
	"auth_service/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract and validate JWT from the request
		token, err := utils.ExtractTokenFromHeader(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// todo array userId and role
		userId, err := utils.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Set user information in the context for further processing
		c.Set("userId", userId)
		c.Next()
	}
}
