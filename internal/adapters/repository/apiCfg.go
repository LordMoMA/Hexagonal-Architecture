package repository

import (
	"errors"
	"os"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/config"
	"github.com/joho/godotenv"
)

func LoadAPIConfig() (*config.APIConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	apiKey := os.Getenv("API_KEY")
	stripeKey := os.Getenv("STRIPE_PRIVATE_KEY")

	if len(jwtSecret) == 0 {
		return nil, errors.New("JWT secret not found")
	}

	return &config.APIConfig{
		JWTSecret: jwtSecret,
		APIKey:    apiKey,
		StripeKey: stripeKey,
	}, nil
}
