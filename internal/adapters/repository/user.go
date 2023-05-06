package repository

import (
	"errors"
	"fmt"
	"os"
	"time"

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

type LoginResponse struct {
	ID 		 string `json:"id"`
	Email 	 string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Membership  bool   `json:"membership"`
}

func (u *DB) LoginUser(email, password string) (*LoginResponse, error) {
    apiCfg, err := loadAPIConfig()
    if err != nil {
        return nil, err
    }

    user, err := u.findUserByEmail(email)
    if err != nil {
        return nil, err
    }

    err = u.verifyPassword(user.Password, password)
    if err != nil {
        return nil, err
    }

    accessToken, err := u.generateAccessToken(user.ID, apiCfg.JWTSecret)
    if err != nil {
        return nil, err
    }

    refreshToken, err := u.generateRefreshToken(user.ID, apiCfg.JWTSecret)
    if err != nil {
        return nil, err
    }

    return &LoginResponse{
        ID:           user.ID,
        Email:        user.Email,
        Membership:   user.Membership,
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }, nil
}

func loadAPIConfig() (*config.APIConfig, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    jwtSecret := os.Getenv("JWT_SECRET")
    apiKey := os.Getenv("API_KEY")

    if len(jwtSecret) == 0 {
        return nil, errors.New("JWT secret not found")
    }

    return &config.APIConfig{
        JWTSecret: jwtSecret,
        APIKey:    apiKey,
    }, nil
}

func (u *DB) findUserByEmail(email string) (*domain.User, error) {
    user := &domain.User{}
    req := u.db.First(&user, "email = ?", email)
    if req.RowsAffected == 0 {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (u *DB) verifyPassword(hash, password string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        return errors.New("password not matched")
    }
    return nil
}

func (u *DB) generateAccessToken(userID, jwtSecret string) (string, error) {
    claims := jwt.RegisteredClaims{
        Issuer:    "LordMoMA",
        Subject:   userID,
        IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour).UTC()),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(jwtSecret))
}

func (u *DB) generateRefreshToken(userID, jwtSecret string) (string, error) {
    claims := jwt.RegisteredClaims{
        Issuer:    "LordMoMA",
        Subject:   userID,
        IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * 24 * time.Hour).UTC()),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(jwtSecret))
}
