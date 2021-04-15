package controllers

import (
	pgModel "github.com/JesusG2000/hexsatisfaction-model/model"
	"github.com/JesusG2000/hexsatisfaction/model"
)

// User is a user service.
type User struct {
	UserDB UserDB
}

// NewUser is a User service constructor.
func NewUser(userDB UserDB) *User {
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
func (u User) IsExist(user pgModel.User) bool {
	return false
}
