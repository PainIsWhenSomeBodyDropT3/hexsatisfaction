package repository

import (
	"database/sql"
	"time"

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

// Purchase is an interface for purchase repository methods.
type Purchase interface {
	Create(purchase model.Purchase) (int, error)
	Delete(id int) (int, error)
	FindById(id int) (*model.Purchase, error)
	FindLastByUserId(id int) (*model.Purchase, error)
	FindAllByUserId(id int) ([]model.Purchase, error)
	FindByUserIdAndPeriod(id int, start, end time.Time) ([]model.Purchase, error)
	FindByUserIdAfterDate(id int, start time.Time) ([]model.Purchase, error)
	FindByUserIdBeforeDate(id int, end time.Time) ([]model.Purchase, error)
	FindByUserIdAndFileName(id int, name string) ([]model.Purchase, error)
	FindLast() (*model.Purchase, error)
	FindAll() ([]model.Purchase, error)
	FindByPeriod(start, end time.Time) ([]model.Purchase, error)
	FindAfterDate(start time.Time) ([]model.Purchase, error)
	FindBeforeDate(end time.Time) ([]model.Purchase, error)
	FindByFileName(name string) ([]model.Purchase, error)
}

// Repositories collects all repository interfaces.
type Repositories struct {
	User     User
	UserRole UserRole
	Purchase Purchase
}

// NewRepositories is a Repositories constructor.
func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User:     NewUserRepo(db),
		UserRole: NewUserRoleRepo(db),
		Purchase: NewPurchaseRepo(db),
	}
}
