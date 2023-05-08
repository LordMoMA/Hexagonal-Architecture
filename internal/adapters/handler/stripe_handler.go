package handler

import (
	"net/http"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type PaymentHandler struct {
	svc services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		svc: paymentService,
	}
}

// CreateCheckoutSessionRequest
type CreatePaymentRequest struct {
	ProductName        string `json:"product_name" binding:"required"`
	ProductDescription string `json:"product_description" binding:"required"`
	Amount             int64  `json:"amount" binding:"required"`
	Currency           string `json:"currency" binding:"required"`
	SuccessURL         string `json:"success_url" binding:"required"`
	CancelURL          string `json:"cancel_url" binding:"required"`
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

	// Call the Stripe checkout API
	params := &stripe.PaymentIntentParams{
		Amount:              stripe.Int64(1099),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes:  []*string{stripe.String("card")},
		StatementDescriptor: stripe.String("Custom descriptor"),
	}
	pi, _ := paymentintent.New(params)

	// Return the session ID
	ctx.JSON(http.StatusOK, gin.H{"client_secret": pi.ClientSecret})
}
