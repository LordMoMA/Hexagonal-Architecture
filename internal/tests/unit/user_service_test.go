package unit

import (
	"testing"
	"time"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/cache"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setUpDB() *repository.DB {
	db, _ := gorm.Open("postgres", "postgres://test:test@localhost:5432/testdb?sslmode=disable")
	db.AutoMigrate(&domain.Message{}, &domain.User{}, &domain.Payment{})

	redisCache, err := cache.NewRedisCache("localhost:6379", "")
	if err != nil {
		panic(err)
	}

	store := repository.NewDB(db, redisCache)

	return store
}

func TestCreateUser(t *testing.T) {
	db := setUpDB()

	email := "alanmoore@example.com"
	password := "password"

	user, err := db.CreateUser(email, password)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
}

func TestReadUser(t *testing.T) {
	db := setUpDB()

	email := "test@example.com"
	password := "password"

	user, err := db.CreateUser(email, password)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	cachedUser, err := db.ReadUser(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, cachedUser)
	assert.Equal(t, user.ID, cachedUser.ID)
	assert.Equal(t, user.Email, cachedUser.Email)
	assert.Equal(t, user.Password, cachedUser.Password)

	time.Sleep(time.Minute * 11)

	cachedUser, err = db.ReadUser(user.ID)
	assert.Error(t, err)
	assert.Nil(t, cachedUser)
}

func TestReadUsers(t *testing.T) {
	db := setUpDB()

	email := "test@example.com"
	password := "password"

	user, err := db.CreateUser(email, password)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	users, err := db.ReadUsers()
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.NotEmpty(t, users)
}

func TestUpdateUser(t *testing.T) {
	db := setUpDB()

	email := "test@example.com"
	password := "password"

	user, err := db.CreateUser(email, password)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	newEmail := "new@example.com"
	newPassword := "newpassword"

	err = db.UpdateUser(user.ID, newEmail, newPassword)
	assert.NoError(t, err)

	cachedUser, err := db.ReadUser(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, cachedUser)
	assert.Equal(t, newEmail, cachedUser.Email)
	assert.NotEqual(t, password, cachedUser.Password)
}

func TestDeleteUser(t *testing.T) {
	db := setUpDB()

	email := "test@example.com"
	password := "password"

	user, err := db.CreateUser(email, password)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	err = db.DeleteUser(user.ID)
	assert.NoError(t, err)

	cachedUser, err := db.ReadUser(user.ID)
	assert.Error(t, err)
	assert.Nil(t, cachedUser)

	users, err := db.ReadUsers()
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Empty(t, users)
}

func TestCreateUserAlreadyExists(t *testing.T) {
	db := setUpDB()

	email := "test@example.com"
	password := "password"

	user, err := db.CreateUser(email, password)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	user, err = db.CreateUser(email, password)
	assert.Error(t, err)
	assert.Nil(t, user)
	// assert.True(t, errors.Is(err, repository.ErrUserAlreadyExists))
}

func TestReadUserNotFound(t *testing.T) {
	db := setUpDB()

	user, err := db.ReadUser("nonexistent")
	assert.Error(t, err)
	assert.Nil(t, user)
	// assert.True(t, errors.Is(err, repository.ErrUserNotFound))
}

func TestUpdateUserNotFound(t *testing.T) {
	db := setUpDB()

	err := db.UpdateUser("nonexistent", "new@example.com", "newpassword")
	assert.Error(t, err)
	// assert.True(t, errors.Is(err, repository.ErrUserNotFound))
}

func TestDeleteUserNotFound(t *testing.T) {
	db := setUpDB()

	err := db.DeleteUser("nonexistent")
	assert.Error(t, err)
	// assert.True(t, errors.Is(err, repository.ErrUserNotFound))
}
