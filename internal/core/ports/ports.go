package ports

import "github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"


type MessengerService interface {
   CreateMessage(message domain.Message) error
   ReadMessage(id string) (*domain.Message, error)
   ReadMessages() ([]*domain.Message, error)
   UpdateMessage(id string, message domain.Message) error
   DeleteMessage(id string) error
}


type MessengerRepository interface {
   CreateMessage(message domain.Message) error
   ReadMessage(id string) (*domain.Message, error)
   ReadMessages() ([]*domain.Message, error)
   UpdateMessage(id string, message domain.Message) error
   DeleteMessage(id string) error
}

type UserRepository interface {
   CreateUser(email, password string) (*domain.User, error)
   ReadUser(id string) (*domain.User, error)
   ReadUsers() ([]*domain.User, error)
   UpdateUser(id string, user domain.User) error
   DeleteUser(id string) error
}


