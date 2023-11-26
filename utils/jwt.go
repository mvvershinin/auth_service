package utils

import (
	"auth_service/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

// todo replace it with repository model
//type User struct {
//	Login    string
//	Password string // For simplicity, store the password in plain text. In a real-world scenario, use hashing.
//}

type TokenClaims struct {
	Login  string `json:"login"`
	UserId int    `json:"userId"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key")

func ExtractTokenFromHeader(r *http.Request) (string, error) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		return "", jwt.ErrInvalidKey
	}

	parts := strings.Split(tokenHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", jwt.ErrInvalidKey
	}

	return parts[1], nil
}

func VerifyToken(tokenString string) (int, error) {
	claims := &TokenClaims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return 0, err
			//jsonResponse(errorDate{Error: "token invalid"}, http.StatusUnauthorized, false, w)
		}
		return 0, err
		//jsonResponse(errorDate{Error: "something wrong"}, http.StatusInternalServerError, false, w)
	}
	if !tkn.Valid {
		return 0, jwt.ErrSignatureInvalid
	}
	return claims.UserId, nil
}

func CreateToken(user models.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = user.Login
	claims["userId"] = user.Id
	//todo role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
