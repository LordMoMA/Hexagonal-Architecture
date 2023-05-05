package ports

import "github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"


type MessengerService interface {
   CreateMessage(message domain.Message) error
   ReadMessage(id string) (*domain.Message, error)
   ReadMessages() ([]*domain.Message, error)
   UpdateMessage(id string) (*domain.Message, error)
   DeleteMessage(id string) error
}


type MessengerRepository interface {
   CreateMessage(message domain.Message) error
   ReadMessage(id string) (*domain.Message, error)
   ReadMessages() ([]*domain.Message, error)
   UpdateMessage(id string) (*domain.Message, error)
   DeleteMessage(id string) error
}


