package handlers

import (
	"auth_service/models"
	"auth_service/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	//	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// todo replace it with repository model
type User struct {
	Login    string
	Password string // For simplicity, store the password in plain text. In a real-world scenario, use hashing.
}

// todo replace it with repository model
//var users = map[string]models.User{
//	"login_string": User{Login: "login_string", Password: "password_string"},
//}

func Login(c *gin.Context) {
	var credentials struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// todo Validate the credentials
	//user, found := users[credentials.Login]
	//if !found || user.Password != credentials.Password {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	//	return
	//}

	user := models.User{
		Login:    "login_string",
		Password: "password_string",
		Id:       434667,
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	//todo write tokenString to redis
	fmt.Printf("%+v\n", token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Verify(c *gin.Context) {
	// todo its userId now, make more info
	userClaims, exists := c.Get("userId")
	if !exists {
		// Handle the case where user claims are not found in the context
		c.JSON(500, gin.H{"error": "User claims not found in context"})
		return
	}
	//claims, _ := userClaims.(utils.TokenClaims)
	//fmt.Println("User ID:", claims.Login)
	//fmt.Println("Audience:", claims.)
	//profile, _ := json.Marshal(claims)
	c.JSON(http.StatusOK, gin.H{"userId": userClaims})
}
