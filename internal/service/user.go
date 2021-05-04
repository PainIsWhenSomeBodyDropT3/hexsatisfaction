package service

import (
	"strconv"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/pkg/errors"
)

// UserService is a user service.
type UserService struct {
	repository.User
	auth.TokenManager
}

// NewUserService is a UserService service constructor.
func NewUserService(userRepo repository.User, tokenManager auth.TokenManager) *UserService {
	return &UserService{userRepo, tokenManager}
}

// Create creates new user and returns id.
func (u UserService) Create(req model.RegisterUserRequest) (int, error) {
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}
	id, err := u.User.Create(user)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't create a user")
	}
	return id, nil
}

// FindByLogin finds the user by login.
func (u UserService) FindByLogin(login string) (*model.User, error) {
	user, err := u.User.FindByLogin(login)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find a user by login")
	}

	return user, nil
}

// FindByCredentials finds the user by credentials and return's jwt-token.
func (u UserService) FindByCredentials(req model.LoginUserRequest) (string, error) {
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}
	newUser, err := u.User.FindByCredentials(user)
	if err != nil {
		return "", errors.Wrap(err, "couldn't find a user by credentials")
	}

	if newUser.ID != 0 {
		newJWT, err := u.NewJWT(strconv.Itoa(newUser.ID))
		if err != nil {
			return "", errors.Wrap(err, "couldn't create a token")
		}
		return newJWT, nil
	}

	return "", nil
}

// IsExist checks if the user exists.
func (u UserService) IsExist(login string) (bool, error) {
	exist, err := u.User.IsExist(login)
	if err != nil {
		return false, errors.Wrap(err, "couldn't check user existence")
	}

	return exist, nil
}
