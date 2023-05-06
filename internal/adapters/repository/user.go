package repository

import (
	"errors"
	"fmt"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/google/uuid"
)


func (u *DB) CreateUser(email, password string) (*domain.User, error) {

 	user := &domain.User{}
	req := u.db.First(&user, "email = ?", email)
	if req.RowsAffected != 0 {
		return nil, errors.New("user already exists")
	}
	
	user = &domain.User{
		ID: uuid.New().String(),
		Email: email,
		Password: password,
		Membership: false,

	}
	req = u.db.Create(&user)
	if req.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("user not saved: %v", req.Error))
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
		return nil, errors.New(fmt.Sprintf("users not found: %v", req.Error))
	}
	return users, nil
}


func (u *DB) UpdateUser(id string, user domain.User) error {
	req := u.db.Model(&user).Where("id = ?", id).Update(user)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
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





