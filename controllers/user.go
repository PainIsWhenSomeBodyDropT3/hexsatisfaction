package controllers

import (
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
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}
	err := u.UserDB.Create(user)
	if err != nil {
		return err
	}

	return nil
}

// FindByLogin find user by login.
func (u User) FindByLogin(login string) (*model.User, error) {
	user, err := u.UserDB.FindByLogin(login)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByCredentials find user by credentials user.
func (u User) FindByCredentials(req model.LoginUserRequest) (*model.User, error) {
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}
	newUser, err := u.UserDB.FindByCredentials(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// IsExist check is  user exist.
func (u User) IsExist(login string) (bool, error) {
	exist, err := u.UserDB.IsExist(login)
	if err != nil {
		return false, err
	}

	return exist, nil
}
