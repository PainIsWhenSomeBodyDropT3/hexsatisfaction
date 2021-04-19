package repository

import "github.com/JesusG2000/hexsatisfaction/model"

// UserRepo is an interface for user repository methods.
type UserRepo interface {
	Create(user model.User) error
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(user model.User) (*model.User, error)
	IsExist(login string) (bool, error)
}
