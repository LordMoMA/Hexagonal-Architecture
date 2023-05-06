package handler

import (
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})

		return
	}

	_, err := h.svc.CreateUser(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) ReadUsers(ctx *gin.Context) {
	
	users, err := h.svc.ReadUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})


	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Token not found",
		})
	}
	tokenString := authHeader[7:]

	// parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{},error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %w", token.Header["alg"])
		}
		return []byte(apiCfg.JWTSecret), nil
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

	if !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Token not valid",
		})
	}

	// check token has expired or not
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || claims.ExpiresAt == nil || claims.ExpiresAt.Before(time.Now().UTC()) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Token has expired",
		})
	}

	// check if token is a refresh token
	if claims.Issuer == "LordMoMA-refresh" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Token is a refresh token, please use access token.",
		})
	}

	id := ctx.Param("id")
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})

		return
	}

	err := h.svc.UpdateUser(id, user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.svc.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func (h *UserHandler) LoginUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})

		return
	}

	response, err := h.svc.LoginUser(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
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