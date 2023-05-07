package handler

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(authHeader string, jwtSecret string) (string, error) {
	// Check if token exists in the header
	if authHeader == "" {
		return "", errors.New("token not found")
	}
	
	// Extract token from header
	tokenString := authHeader[7:]

	// Parse and validate token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token not valid")
	}

	// Check if token has expired
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || claims.ExpiresAt == nil || claims.ExpiresAt.Before(time.Now().UTC()) {
		return "", errors.New("token has expired")
	}

	// Check if token is a refresh token
	if claims.Issuer == "LordMoMA-refresh" {
		return "", errors.New("token is a refresh token, please use access token")
	}

	// Extract user ID from token
	userID := claims.Subject

	return userID, nil
}