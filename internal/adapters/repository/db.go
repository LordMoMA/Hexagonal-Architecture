package repository

import (
	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/cache"
	"github.com/jinzhu/gorm"
)

type DB struct {
	db    *gorm.DB
	cache *cache.RedisCache
}

// new database
func NewDB(db *gorm.DB, cache *cache.RedisCache) *DB {
	return &DB{
		db:    db,
		cache: cache,
	}
}
