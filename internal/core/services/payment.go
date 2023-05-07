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
