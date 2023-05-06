package handler

import "github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(UserService services.UserService) *UserHandler {
	return &UserHandler{
		svc: UserService,
	}
}