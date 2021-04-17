package repository

import "github.com/JesusG2000/hexsatisfaction/model"

// UserDB is an interface for UserRepo methods.
type UserDB interface {
	Create(user model.User) error
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(user model.User) (*model.User, error)
	IsExist(login string) (bool, error)
}
