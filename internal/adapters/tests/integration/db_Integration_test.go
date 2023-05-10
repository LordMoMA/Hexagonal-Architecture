package integration

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/cache"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func TestDBIntegration(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbPassword, dbname)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	redisCache, err := cache.NewRedisCache("localhost:6379", "")
	if err != nil {
		panic(err)
	}
	
	store := repository.NewDB(db, redisCache)

	// create a test user
	email := "test@example.com"
	password := "password"
	user, err := store.CreateUser(email, password)
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	// test reading a user
	readUser, err := store.ReadUser(user.ID)
	if err != nil {
		t.Fatalf("failed to read user: %v", err)
	}
	if readUser.Email != email {
		t.Errorf("expected email %q, got %q", email, readUser.Email)
	}

	// test reading all users
	users, err := store.ReadUsers()
	if err != nil {
		t.Fatalf("failed to read users: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("expected 1 user, got %d", len(users))
	}
	if users[0].Email != email {
		t.Errorf("expected email %q, got %q", email, users[0].Email)
	}

	// test updating a user
	newEmail := "newemail@example.com"
	newPassword := "newpassword"
	err = store.UpdateUser(user.ID, newEmail, newPassword)
	if err != nil {
		t.Fatalf("failed to update user: %v", err)
	}
	readUser, err = store.ReadUser(user.ID)
	if err != nil {
		t.Fatalf("failed to read updated user: %v", err)
	}
	if readUser.Email != newEmail {
		t.Errorf("expected email %q, got %q", newEmail, readUser.Email)
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Errorf("password not hashed: %v", err)
	}
	
	if readUser.Password != string(hashedNewPassword) {
		t.Errorf("expected password %q, got %q", newPassword, readUser.Password)
	}

	// test deleting a user
	err = store.DeleteUser(user.ID)
	if err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}
	_, err = store.ReadUser(user.ID)
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("expected user not found error, got %v", err)
	}
}
