package integration

import (
	"testing"
	"time"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/rogpeppe/go-internal/cache"
)

func BenchmarkCache(b *testing.B, cache cache.Cache) {
	user := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}
	key := "testKey"

	// Set the value in the cache
	err := cache.Set(key, user, 10*time.Minute)
	if err != nil {
		b.Fatalf("error setting value in cache: %v", err)
	}

	// Measure the time taken to retrieve the value from the cache
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result domain.User
		err := cache.Get(key, &result)
		if err != nil {
			b.Fatalf("error getting value from cache: %v", err)
		}
	}
}
