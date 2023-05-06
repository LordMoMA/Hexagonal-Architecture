package handler

import (
	"net/http"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"
	"github.com/gin-gonic/gin"
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