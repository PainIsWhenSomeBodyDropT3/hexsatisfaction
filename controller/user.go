package controller

import (
	"github.com/JesusG2000/hexsatisfaction/model"
	"github.com/pkg/errors"
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
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}
	err := u.UserDB.Create(user)
	if err != nil {
		return errors.Wrap(err, "couldn't create a user.")
	}

	return nil
}

// FindByLogin finds the user by login.
func (u User) FindByLogin(login string) (*model.User, error) {
	user, err := u.UserDB.FindByLogin(login)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find a user by login.")
	}

	return user, nil
}

// FindByCredentials finds the user by credentials.
func (u User) FindByCredentials(req model.LoginUserRequest) (*model.User, error) {
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}
	newUser, err := u.UserDB.FindByCredentials(user)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find a user by credentials.")
	}

	return newUser, nil
}

// IsExist checks if the user exists.
func (u User) IsExist(login string) (bool, error) {
	exist, err := u.UserDB.IsExist(login)
	if err != nil {
		return false, errors.Wrap(err, "user not exist errors.")
	}

	return exist, nil
}
