package repository

import "github.com/LordMoMA/Hexagonal-Architecture/internal/core/ports"

type PaymentService struct {
	repo ports.PaymentRepository
}
