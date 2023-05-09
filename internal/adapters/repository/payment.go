package repository

import (
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/google/uuid"
)

// deposit
func (p *DB) Deposit(userID string, amount string, payment domain.Payment) error {
	payment.OrderID = uuid.New().String()
	payment = domain.Payment{
		OrderID: payment.OrderID,
		UserID:  userID,
		Amount:  payment.Amount,
		Status: "success",
	}
	req := p.db.Create(&payment)
	if req.RowsAffected == 0 {
		return req.Error
	}
	return nil
}
// withdraw
func (p *DB) Withdraw(userID string, amount string, payment domain.Payment) error {
	payment.OrderID = uuid.New().String()
	payment = domain.Payment{
		OrderID: payment.OrderID,	
		UserID:  userID,
		Amount:  payment.Amount,
		Status: "success",
	}
	req := p.db.Create(&payment)
	if req.RowsAffected == 0 {
		return req.Error
	}
	return nil
}

// create payment
func (p *DB) ProcessPaymentWithStripe(userID string, payment domain.Payment) error {
	payment.OrderID = uuid.New().String()
	payment = domain.Payment{
		OrderID: payment.OrderID,
		UserID:  userID,
		Amount:  payment.Amount,
	}
	req := p.db.Create(&payment)
	if req.RowsAffected == 0 {
		return req.Error
	}
	return nil
}

