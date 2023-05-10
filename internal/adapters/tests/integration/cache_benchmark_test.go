package integration

import (
	"testing"
	"time"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/ports"
)

func BenchmarkCache(b *testing.B) {
	user := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}
	key := "testKey"

	// Set the value in the cache
	err := ports.CacheRepository.Set(key, user, 10*time.Minute)
	if err != nil {
		b.Fatalf("error setting value in cache: %v", err)
	}

	// Measure the time taken to retrieve the value from the cache
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result domain.User
		err := ports.CacheRepository.Get(key, &result)
		if err != nil {
			b.Fatalf("error getting value from cache: %v", err)
		}
	}
}
