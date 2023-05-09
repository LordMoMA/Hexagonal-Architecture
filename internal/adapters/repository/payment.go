package repository

import (
	"fmt"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
)

//get balance
func (p *DB) GetBalance(userID string) (string, error) {
	var balance string
	req := p.db.Table("payments").Where("user_id = ?", userID).Select("sum(amount) as balance").Row()
	err := req.Scan(&balance)
	if err != nil {
		return "", err
	}
	return balance, nil
}

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

func getOrderIDFromStripeSession(sessionID string) (string, error) {
	apiCfg, err := LoadAPIConfig()
	if err != nil {
		return "", err
	}
	stripe.Key = apiCfg.StripeKey

    // Retrieve the Stripe Checkout Session from the API
    checkoutSession, err := session.Get(sessionID, nil)

    if err != nil {
        return "", fmt.Errorf("unable to retrieve checkout session from Stripe: %v", err)
    }

    // Retrieve the order ID from the session metadata
    orderID, ok := checkoutSession.Metadata["order_id"]

    if !ok {
        return "", fmt.Errorf("order ID not found in session metadata")
    }

    return orderID, nil
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

