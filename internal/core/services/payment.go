package services

import "github.com/LordMoMA/Hexagonal-Architecture/internal/core/ports"

type PaymentService struct {
	repo ports.PaymentRepository
}

func NewPaymentService(repo ports.PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

func (u *UserService) UpdateMembershipStatus(id string, status bool) error {
	return u.repo.UpdateMembershipStatus(id, status)
}
