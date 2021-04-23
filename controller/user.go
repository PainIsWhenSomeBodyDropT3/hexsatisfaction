package controller

import (
	e "github.com/JesusG2000/hexsatisfaction/errors"
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
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}
	err := u.UserDB.Create(user)
	if err != nil {
		return e.DatabaseError("couldn't create a user.", err)
	}

	return nil
}

// FindByLogin find user by login.
func (u User) FindByLogin(login string) (*model.User, error) {
	user, err := u.UserDB.FindByLogin(login)
	if err != nil {
		return nil, e.DatabaseError("couldn't find a user by login.", err)
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
		return nil, e.DatabaseError("couldn't find a user by credentials.", err)
	}

	return newUser, nil
}

// IsExist check is  user exist.
func (u User) IsExist(login string) (bool, error) {
	exist, err := u.UserDB.IsExist(login)
	if err != nil {
		return false, e.DatabaseError("user not exist errors.", err)
	}

	return exist, nil
}
