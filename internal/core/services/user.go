package services

import "github.com/LordMoMA/Hexagonal-Architecture/internal/core/ports"

type UserService struct {
	repo ports.UserRepository
}