package ports

import (
	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
)

type MessengerService interface {
	CreateMessage(userID string, message domain.Message) error
	ReadMessage(id string) (*domain.Message, error)
	ReadMessages() ([]*domain.Message, error)
	UpdateMessage(id string, message domain.Message) error
	DeleteMessage(id string) error
}

type MessengerRepository interface {
	CreateMessage(userID string, message domain.Message) error
	ReadMessage(id string) (*domain.Message, error)
	ReadMessages() ([]*domain.Message, error)
	UpdateMessage(id string, message domain.Message) error
	DeleteMessage(id string) error
}

type UserService interface {
	CreateUser(email, password string) (*domain.User, error)
	ReadUser(id string) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	UpdateUser(id, email, password string) error
	DeleteUser(id string) error
	LoginUser(email, password string) (*repository.LoginResponse, error)
	UpdateMembershipStatus(id string, status bool) error
}

type UserRepository interface {
	CreateUser(email, password string) (*domain.User, error)
	ReadUser(id string) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	UpdateUser(id, email, password string) error
	DeleteUser(id string) error
	LoginUser(email, password string) (*repository.LoginResponse, error)
	UpdateMembershipStatus(id string, status bool) error
}

type PaymentService interface {
	ProcessPaymentWithStripe(userID string, payment domain.Payment) error
}

type PaymentRepository interface {
	ProcessPaymentWithStripe(userID string, payment domain.Payment) error
}
