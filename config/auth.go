package config

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecretkey")

// GenerateToken creates a JWT token for authentication
func GenerateToken(userID uint) (string, error) {
	claims := &jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken checks if a JWT is valid
func ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
