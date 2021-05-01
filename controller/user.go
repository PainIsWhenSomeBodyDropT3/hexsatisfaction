package controller

import (
	"github.com/JesusG2000/hexsatisfaction/jwt"
	"github.com/JesusG2000/hexsatisfaction/model"
	"github.com/JesusG2000/hexsatisfaction/repository"
	"github.com/pkg/errors"
)

// User is a user service.
type User struct {
	repository.UserRepo
	*jwt.Manager
}

// NewUser is a User service constructor.
func NewUser(userRepo repository.UserRepo, tokenManager *jwt.Manager) *User {
	return &User{userRepo, tokenManager}
}

// Create creates new user.
func (u User) Create(req model.RegisterUserRequest) error {
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}
	err := u.UserRepo.Create(user)
	if err != nil {
		return errors.Wrap(err, "couldn't create a user")
	}
	return nil
}

// FindByLogin finds the user by login.
func (u User) FindByLogin(login string) (*model.User, error) {
	user, err := u.UserRepo.FindByLogin(login)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find a user by login")
	}

	return user, nil
}

// FindByCredentials finds the user by credentials and return's jwt-token.
func (u User) FindByCredentials(req model.LoginUserRequest) (string, error) {
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}
	newUser, err := u.UserRepo.FindByCredentials(user)
	if err != nil {
		return "", errors.Wrap(err, "couldn't find a user by credentials")
	}

	if newUser.ID != 0 {
		newJWT, err := u.NewJWT(string(rune(newUser.ID)))
		if err != nil {
			return "", errors.Wrap(err, "couldn't create a token")
		}
		return newJWT, nil
	}

	return "", nil
}

// IsExist checks if the user exists.
func (u User) IsExist(login string) (bool, error) {
	exist, err := u.UserRepo.IsExist(login)
	if err != nil {
		return false, errors.Wrap(err, "couldn't check user existence")
	}

	return exist, nil
}
