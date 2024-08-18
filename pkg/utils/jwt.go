package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/priyansurout/auth-microservice/internal/user"
)

var jwtKey = []byte("your-secret-key") // Replace with a secure key

func GenerateJWT(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Set token expiration

	claims := &user.Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
