package services

import (
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/ports"
	"github.com/google/uuid"
)


type MessengerService struct {
   repo ports.MessengerRepository
}


func NewMessengerService(repo ports.MessengerRepository) *MessengerService {
   return &MessengerService{
       repo: repo,
   }
}


func (m *MessengerService) SaveMessage(message domain.Message) error {
   message.ID = uuid.New().String()
   return m.repo.SaveMessage(message)
}


func (m *MessengerService) ReadMessage(id string) (*domain.Message, error) {
   return m.repo.ReadMessage(id)
}


func (m *MessengerService) ReadMessages() ([]*domain.Message, error) {
   return m.repo.ReadMessages()
}