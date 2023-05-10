package integration

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func TestDBIntegration(t *testing.T) {
	// initialize the database and cache
	
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
	db := repository.NewDB(db, redisCache)

	// create a test user
	email := "test@example.com"
	password := "password"
	user, err := db.CreateUser(email, password)
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	// test reading a user
	readUser, err := db.ReadUser(user.ID)
	if err != nil {
		t.Fatalf("failed to read user: %v", err)
	}
	if readUser.Email != email {
		t.Errorf("expected email %q, got %q", email, readUser.Email)
	}

	// test reading all users
	users, err := db.ReadUsers()
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
	err = db.UpdateUser(user.ID, newEmail, newPassword)
	if err != nil {
		t.Fatalf("failed to update user: %v", err)
	}
	readUser, err = db.ReadUser(user.ID)
	if err != nil {
		t.Fatalf("failed to read updated user: %v", err)
	}
	if readUser.Email != newEmail {
		t.Errorf("expected email %q, got %q", newEmail, readUser.Email)
	}
	if !domain.CheckPassword(newPassword, readUser.Password) {
		t.Errorf("password not updated")
	}

	// test deleting a user
	err = db.DeleteUser(user.ID)
	if err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}
	_, err = db.ReadUser(user.ID)
	if !errors.Is(err, database.ErrUserNotFound) {
		t.Errorf("expected user not found error, got %v", err)
	}
}
