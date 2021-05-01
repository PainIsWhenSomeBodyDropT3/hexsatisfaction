package handler

import (
	"github.com/JesusG2000/hexsatisfaction/model"
)

// UserService is an interface for User methods.
type UserService interface {
	Create(req model.RegisterUserRequest) error
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(req model.LoginUserRequest) (string, error)
	IsExist(login string) (bool, error)
}
