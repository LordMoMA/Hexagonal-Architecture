package repository

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/config"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


func (u *DB) CreateUser(email, password string) (*domain.User, error) {

 	user := &domain.User{}
	req := u.db.First(&user, "email = ?", email)
	if req.RowsAffected != 0 {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("password not hashed: %v", err)
	}
	
	user = &domain.User{
		ID: uuid.New().String(),
		Email: email,
		Password: string(hashedPassword),
		Membership: false,
	}
	req = u.db.Create(&user)
	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("user not saved: %v", req.Error)
	}
	return user, nil
}


func (u *DB) ReadUser(id string) (*domain.User, error) {
	user := &domain.User{}
	req := u.db.First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}


func (u *DB) ReadUsers() ([]*domain.User, error) {
	var users []*domain.User
	req := u.db.Find(&users)
	if req.Error != nil {
		return nil, fmt.Errorf("users not found: %v", req.Error)
	}
	return users, nil
}


func (u *DB) UpdateUser(id, email, password string) error {
	// get user by id
	user := &domain.User{}
	req := u.db.First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("password not hashed: %v", err)
	}

	user = &domain.User{
		Email: email,
		Password: string(hashedPassword),
	}
	req = u.db.Model(&user).Where("id = ?", id).Update(user)
	if req.RowsAffected == 0 {
		return errors.New("unable to update user :(")
	}
	return nil

}


func (u *DB) DeleteUser(id string) error {
	user := &domain.User{}
	req := u.db.Where("id = ?", id).Delete(&user)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}



func (u *DB) LoginUser(email, password string) (*domain.User, error) {
	
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

   jwtSecret := os.Getenv("JWT_SECRET")
   apiKey := os.Getenv("API_KEY")

   apiCfg := &config.APIConfig{
      JWTSecret: jwtSecret,
      APIKey:    apiKey,
   }

	user := &domain.User{}
	req := u.db.First(&user, "email = ?", email)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password not matched: %v", err)
	}
	
	if len(apiCfg.SecretKey) == 0 {
		return nil, errors.New("secret key not found")
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"password": password,
	})


	return user, nil
}




