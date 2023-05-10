package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Membership   bool   `json:"membership"`
}

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
		ID:         uuid.New().String(),
		Email:      email,
		Password:   string(hashedPassword),
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
	cachekey := user.ID
	err := u.cache.Get(cachekey, &user)
	if err == nil {
		return user, nil
	}

	req := u.db.First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	err = u.cache.Set(cachekey, user, time.Minute*10)
	if err != nil {
		fmt.Printf("Error storing user in cache: %v", err)
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
		Email:    email,
		Password: string(hashedPassword),
	}
	req = u.db.Model(&user).Where("id = ?", id).Update(user)
	if req.RowsAffected == 0 {
		return errors.New("unable to update user :(")
	}

	// delete user in the cache
	err = u.cache.Delete(id)
	if err != nil {
		fmt.Printf("Error deleting user in cache: %v", err)
	}

	return nil

}

func (u *DB) DeleteUser(id string) error {
	user := &domain.User{}
	req := u.db.Where("id = ?", id).Delete(&user)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}
	err := u.cache.Delete(id)
	if err != nil {
		fmt.Printf("Error deleting user in cache: %v", err)
	}
	return nil
}

func (u *DB) LoginUser(email, password string) (*LoginResponse, error) {
	apiCfg, err := LoadAPIConfig()
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
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Membership:   user.Membership,
	}, nil
}

func (u *DB) UpdateMembershipStatus(id string, membership bool) error {
	user := &domain.User{}
	req := u.db.First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}

	user = &domain.User{
		Membership: membership,
	}
	req = u.db.Model(&user).Where("id = ?", id).Update(user)
	if req.RowsAffected == 0 {
		return errors.New("unable to update membership status :(")
	}
	return nil
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
		Issuer:    "LordMoMA-access",
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour).UTC()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func (u *DB) generateRefreshToken(userID, jwtSecret string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "LordMoMA-refresh",
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour).UTC()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
