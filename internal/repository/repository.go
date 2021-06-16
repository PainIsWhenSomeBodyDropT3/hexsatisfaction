package repository

import (
	"database/sql"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
)

// User is an interface for UserRepo methods.
type User interface {
	Create(user model.User) (int, error)
	FindByID(id int) (*model.User, error)
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(user model.User) (*model.User, error)
	IsExist(login string) (bool, error)
	IsExistByID(id int) (bool, error)
}

// UserRole is an interface for UserRoleRepo methods.
type UserRole interface {
	FindAllUser() ([]model.User, error)
}

// Author is an interface for AuthorRepo methods.
type Author interface {
	Create(author model.Author) (int, error)
	Update(id int, author model.Author) (int, error)
	Delete(id int) (int, error)
	FindByID(id int) (*model.Author, error)
	IsExistByID(id int) (bool, error)
	FindByUserID(id int) (*model.Author, error)
	FindByName(name string) ([]model.Author, error)
	FindAll() ([]model.Author, error)
}

// Repositories collects all repository interfaces.
type Repositories struct {
	User     User
	UserRole UserRole
	Author   Author
}

// NewRepositories is a Repositories constructor.
func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User:     NewUserRepo(db),
		UserRole: NewUserRoleRepo(db),
		Author:   NewAuthorRepo(db),
	}
}
