package handler

import (
	"log"
	"net/http"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

type PaymentHandler struct {
	svc services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		svc: paymentService,
	}
}

func (h *PaymentHandler) CreateCheckoutSession(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	stripe.Key = apiCfg.StripeKey

	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String("price_1N5VNbKb78q3bJ6obePPkame"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?success=true"),
		CancelURL:  stripe.String(domain + "?canceled=true"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	ctx.Redirect(http.StatusSeeOther, s.URL)

}

/*

// CreateCheckoutSessionRequest
type CreatePaymentRequest struct {
	ProductName        string `json:"product_name" binding:"required"`
	ProductDescription string `json:"product_description" binding:"required"`
	Amount             string `json:"amount" binding:"required"`
	Currency           string `json:"currency" binding:"required"`
	// SuccessURL         string `json:"success_url" binding:"required"`
	// CancelURL          string `json:"cancel_url" binding:"required"`
}

func (h *PaymentHandler) ProcessPaymentWithStripe(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	stripe.Key = apiCfg.StripeKey
	// Parse request parameters
	var req CreatePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Stripe API to create a PaymentIntent
	params := &stripe.PaymentIntentParams{
		Amount:              stripe.Int64(1099),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes:  []*string{stripe.String("card")},
		StatementDescriptor: stripe.String("Custom descriptor"),
	}
	pi, _ := paymentintent.New(params)

	// userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
	// if err != nil {
	// 	HandleError(ctx, http.StatusBadRequest, err)
	// 	return
	// }

	// // Create Payment object in database
	// payment := &domain.Payment{
	// 	OrderID:  pi.ID,
	// 	UserID:   userID,
	// 	Amount:   req.Amount,
	// 	Currency: req.Currency,
	// 	Status:   "pending",
	// }
	// err = h.svc.ProcessPaymentWithStripe(userID, payment)

	// Return client_secret to client
	ctx.JSON(http.StatusOK, gin.H{"client_secret": pi.ClientSecret})
}

*/
