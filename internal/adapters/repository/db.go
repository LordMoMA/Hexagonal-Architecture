package repository

import (
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/jinzhu/gorm"
)

type DB struct {
	db *gorm.DB
}

type DBStructure struct {
	Users map[int]domain.User `json:"users"`
	Messages map[int]domain.Message `json:"messages"`
}

// new database
func NewDB(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}
