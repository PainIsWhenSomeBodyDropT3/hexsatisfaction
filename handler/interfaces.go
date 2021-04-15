package handler

import (
	pgModel "github.com/JesusG2000/hexsatisfaction-model/model"
	"github.com/JesusG2000/hexsatisfaction/model"
)

// UserService is an interface for User methods.
type UserService interface {
	Create(req model.RegisterUserRequest) error
	FindByLogin(login string) (*pgModel.User, error)
	FindByCredentials(req model.LoginUserRequest) (*pgModel.User, error)
	IsExist(login string) (bool, error)
}
