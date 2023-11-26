package middlewares

import (
	"auth_service/models"
	"auth_service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMiddlewareValidToken(t *testing.T) {
	// Create a mock Gin context with a valid JWT token in the Authorization header
	user := models.User{
		Login:    "login_string",
		Password: "password_string",
		Id:       12345,
	}
	token, _ := utils.CreateToken(user)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set up the middleware
	middleware := AuthMiddleware()

	// Call the middleware
	middleware(c)

	// Check that the user information is set in the context
	userId, exists := c.Get("userId")
	assert.True(t, exists)
	assert.Equal(t, 12345, userId)
}

func TestAuthMiddlewareInvalidToken(t *testing.T) {
	// Create a mock Gin context with an invalid JWT token in the Authorization header
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer invalid-jwt-token")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set up the middleware
	middleware := AuthMiddleware()

	// Call the middleware
	middleware(c)

	// Check that the middleware aborts with a 401 Unauthorized status
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
	assert.Empty(t, c.Keys)
}

func TestAuthMiddlewareMissingToken(t *testing.T) {
	// Create a mock Gin context without the Authorization header
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set up the middleware
	middleware := AuthMiddleware()

	// Call the middleware
	middleware(c)

	// Check that the middleware aborts with a 401 Unauthorized status
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
	assert.Empty(t, c.Keys)
}
