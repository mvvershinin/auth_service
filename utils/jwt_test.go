package utils

import (
	"auth_service/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateAndVerifyToken(t *testing.T) {
	// Create a user for testing
	user := models.User{
		Login:    "testuser",
		Password: "testpassword",
		// Add other user details if needed
	}

	// Create a token
	token, err := CreateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify the token
	userID, err := VerifyToken(token)
	assert.NoError(t, err)
	assert.NotZero(t, userID)
}

func TestExtractTokenFromHeader(t *testing.T) {
	// Create a mock HTTP request with the Authorization header
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer your-mock-token")

	// Extract the token from the request
	token, err := ExtractTokenFromHeader(req)
	assert.NoError(t, err)
	assert.Equal(t, "your-mock-token", token)
}

// todo func TestExpiredToken(t *testing.T) {

func TestInvalidToken(t *testing.T) {
	// Verify an invalid token
	userID, err := VerifyToken("invalid-token")
	assert.Error(t, err)
	assert.Zero(t, userID)
}

func TestExtractTokenFromHeaderInvalidFormat(t *testing.T) {
	// Create a mock HTTP request with an invalid Authorization header format
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "InvalidFormatToken")

	// Extract the token from the request with an invalid format
	token, err := ExtractTokenFromHeader(req)
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestExtractTokenFromHeaderMissingToken(t *testing.T) {
	// Create a mock HTTP request without the Authorization header
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	// Extract the token from the request without the Authorization header
	token, err := ExtractTokenFromHeader(req)
	assert.Error(t, err)
	assert.Empty(t, token)
	assert.True(t, errors.Is(err, jwt.ErrInvalidKey))
}
