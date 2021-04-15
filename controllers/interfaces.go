package controllers

import "github.com/JesusG2000/hexsatisfaction-model/model"

// UserDB is an interface for UserRepo methods.
type UserDB interface {
	Create(user model.User) error
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(user model.User) (*model.User, error)
	IsExist(user model.User) bool
}
