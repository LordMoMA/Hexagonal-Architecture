package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhook"
)

type WebhookRequest struct {
	Event  string `json:"event"`
	UserId string `json:"user_id"`
}

func (h *UserHandler) UpdateMembershipStatus(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	// get api key from config
	apiKey := apiCfg.APIKey

	// check if api key is valid
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		HandleError(ctx, http.StatusBadRequest, errors.New("no api key provided"))
		return
	}
	apiString := strings.TrimPrefix(authHeader, "ApiKey ")

	if apiString != apiKey {
		HandleError(ctx, http.StatusBadRequest, errors.New("invalid api key"))
		return
	}

	var req WebhookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	if req.Event != "membership_status_updated" {
		HandleError(ctx, http.StatusBadRequest, errors.New("invalid event type"))
		return
	}

	userId := req.UserId

	err = h.svc.UpdateMembershipStatus(userId, true)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User's membership status updated successfully",
	})
}

// stripe webhook
func handleWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	payload, err := c.GetRawData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook error while parsing basic request. %v\n", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Replace this endpoint secret with your endpoint's unique secret
	// If you are testing with the CLI, find the secret by running 'stripe listen'
	// If you are using an endpoint defined with the API or dashboard, look in your webhook settings
	// at https://dashboard.stripe.com/webhooks
	endpointSecret := "whsec_"
	signatureHeader := c.GetHeader("Stripe-Signature")
	event, err = webhook.ConstructEvent(payload, signatureHeader, endpointSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook signature verification failed. %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}
	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		log.Printf("Successful payment for %d.", paymentIntent.Amount)
		// Then define and call a func to handle the successful payment intent.
		// handlePaymentIntentSucceeded(paymentIntent)
	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handlePaymentMethodAttached(paymentMethod)
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	c.Status(http.StatusOK)
}
