package controllers

import (
	pgModel "github.com/JesusG2000/hexsatisfaction-model/model"
	"github.com/JesusG2000/hexsatisfaction/model"
	"github.com/JesusG2000/hexsatisfaction/repository"
)

// User is c user service.
type User struct {
	UserDB repository.UserDB
}

// NewUser is c User service constructor.
func NewUser(userDB repository.UserDB) *User {
	return &User{userDB}
}

// Create creates new user.
func (u User) Create(req model.RegisterUserRequest) error {
	return nil
}

// FindByLogin find user by login.
func (u User) FindByLogin(login string) (*pgModel.User, error) {
	return nil, nil
}

// FindByCredentials find user by credentials user.
func (u User) FindByCredentials(req model.LoginUserRequest) (*pgModel.User, error) {
	return nil, nil
}

// IsExist check is  user exist.
func (u User) IsExist(login string) (bool, error) {
	return false, nil
}
