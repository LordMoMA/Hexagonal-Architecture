package services

import (
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/ports"
)

type PaymentService struct {
	repo ports.PaymentRepository
}

func NewPaymentService(repo ports.PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

func (p *PaymentService) CreatePayment(userID string, payment domain.Payment) error {
	return p.repo.CreatePayment(userID, payment)
}
