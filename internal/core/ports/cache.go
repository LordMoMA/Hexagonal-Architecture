package ports

import "time"

type CacheRepository interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, value interface{}) error
	Delete(key string) error
}
