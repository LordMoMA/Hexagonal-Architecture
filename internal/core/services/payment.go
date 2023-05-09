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

func (p *PaymentService) ProcessPaymentWithStripe(userID string, payment domain.Payment) error {
	return p.repo.ProcessPaymentWithStripe(userID, payment)
}

// // deposit
// func (p *PaymentService) Deposit(userID string, amount float64) error {
// 	return p.repo.Deposit(userID, amount)
// }

// // withdraw
// func (p *PaymentService) Withdraw(userID string, amount float64) error {
// 	return p.repo.Withdraw(userID, amount)
// }
