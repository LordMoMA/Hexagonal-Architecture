package repository

import (
	"errors"
	"fmt"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// 	type MessengerPostgresRepository struct {
// 	db *gorm.DB
// }


func (m *DB) CreateMessage(message domain.Message) error {
	req := m.db.Create(&message)
	if req.RowsAffected == 0 {
		return fmt.Errorf("messages not saved: %v", req.Error)
	}
	return nil
}

func (m *DB) ReadMessage(id string) (*domain.Message, error) {
	message := &domain.Message{}
	req := m.db.First(&message, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("message not found")
	}
	return message, nil
}

func (m *DB) ReadMessages() ([]*domain.Message, error) {
	var messages []*domain.Message
	req := m.db.Find(&messages)
	if req.Error != nil {
		return nil, fmt.Errorf("messages not found: %v", req.Error)
	}
	return messages, nil
}

func (m *DB) UpdateMessage(id string, message domain.Message) error {
	req := m.db.Model(&message).Where("id = ?", id).Update(message)
	if req.RowsAffected == 0 {
		return errors.New("message not found")
	}
	return nil
}

func (m *DB) DeleteMessage(id string) error {
	message := &domain.Message{}
	req := m.db.Where("id = ?", id).Delete(&message)
	if req.RowsAffected == 0 {
		return errors.New("message not found")
	}
	return nil
}


	
