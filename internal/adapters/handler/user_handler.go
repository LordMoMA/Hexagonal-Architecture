package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(UserService services.UserService) *UserHandler {
	return &UserHandler{
		svc: UserService,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err := h.svc.CreateUser(user.Email, user.Password)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New user created successfully",
	})	
}

func (h *UserHandler) ReadUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.svc.ReadUser(id)

	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) ReadUsers(ctx *gin.Context) {
	
	users, err := h.svc.ReadUsers()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	// Load API configuration
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	// Validate token
	userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	// Update user
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	err = h.svc.UpdateUser(userID, user.Email, user.Password)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	err = h.svc.DeleteUser(userID)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

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


func (h *UserHandler) LoginUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	response, err := h.svc.LoginUser(user.Email, user.Password)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
	"id":             response.ID,
    "email":          response.Email,
    "access_token":   response.AccessToken,
    "refresh_token":  response.RefreshToken,
    "is_member":      response.Membership,
	})
}