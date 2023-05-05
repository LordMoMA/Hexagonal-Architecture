package repository

import (
	"errors"
	"fmt"
	"os"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type MessengerPostgresRepository struct {
	db *gorm.DB
}

// be sure to hide your password in a .env file, this is for simplicity here.
func NewMessengerPostgresRepository() *MessengerPostgresRepository {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.Message{})

	return &MessengerPostgresRepository{
		db: db,
	}
}

func (m *MessengerPostgresRepository) CreateMessage(message domain.Message) error {
	req := m.db.Create(&message)
	if req.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("messages not saved: %v", req.Error))
	}
	return nil
}

func (m *MessengerPostgresRepository) ReadMessage(id string) (*domain.Message, error) {
	message := &domain.Message{}
	req := m.db.First(&message, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("message not found")
	}
	return message, nil
}

func (m *MessengerPostgresRepository) ReadMessages() ([]*domain.Message, error) {
	var messages []*domain.Message
	req := m.db.Find(&messages)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("messages not found: %v", req.Error))
	}
	return messages, nil
}

func (m *MessengerPostgresRepository) UpdateMessage(id string, message domain.Message) error {
	req := m.db.Model(&message).Where("id = ?", id).Update(message)
	if req.RowsAffected == 0 {
		return errors.New("message not found")
	}
	return nil
}

func (m *MessengerPostgresRepository) DeleteMessage(id string) error {
	message := &domain.Message{}
	req := m.db.Where("id = ?", id).Delete(&message)
	if req.RowsAffected == 0 {
		return errors.New("message not found")
	}
	return nil
}


	
