package services

type PaymentService struct {
	repo ports.PaymentRepository
}

func NewPaymentService(repo ports.PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}
