package repository

import (
	"database/sql"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
)

// User is an interface for user repository methods.
type User interface {
	Create(user model.User) (int, error)
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(user model.User) (*model.User, error)
	IsExist(login string) (bool, error)
}

// UserRole is an interface for user role repository methods.
type UserRole interface {
	FindAllUser() ([]model.User, error)
}

// Repositories collects all repository interfaces.
type Repositories struct {
	User     User
	UserRole UserRole
}

// NewRepositories is a Repositories constructor.
func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User:     NewUserRepo(db),
		UserRole: NewUserRoleRepo(db),
	}
}
