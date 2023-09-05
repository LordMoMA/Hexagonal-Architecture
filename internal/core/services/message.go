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

func (m *MessengerService) CreateMessage(userID string, message domain.Message) error {
	message.ID = uuid.New().String()
	return m.repo.CreateMessage(userID, message)
}

func (m *MessengerService) ReadMessage(id string) (*domain.Message, error) {
	return m.repo.ReadMessage(id)
}

func (m *MessengerService) ReadMessages() ([]*domain.Message, error) {
	return m.repo.ReadMessages()
}

func (m *MessengerService) UpdateMessage(id string, message domain.Message) error {
	return m.repo.UpdateMessage(id, message)
}

func (m *MessengerService) DeleteMessage(id string) error {
	return m.repo.DeleteMessage(id)
}
