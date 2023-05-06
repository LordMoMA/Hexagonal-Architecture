package services

import (
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(user domain.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserService) ReadUser(id string) (*domain.User, error) {
	return u.repo.ReadUser(id)
}

func (u *UserService) ReadUsers() ([]*domain.User, error) {
	return u.repo.ReadUsers()
}

func (u *UserService) UpdateUser(id string, user domain.User) error {
	return u.repo.UpdateUser(id, user)
}

func (u *UserService) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}




