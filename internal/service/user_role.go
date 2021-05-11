package service

import (
	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/pkg/errors"
)

// UserRoleService is a user role service.
type UserRoleService struct {
	repository.UserRole
}

// NewUserRoleService is a UserRoleService constructor.
func NewUserRoleService(userRoleRepo repository.UserRole) *UserRoleService {
	return &UserRoleService{userRoleRepo}
}

// FindAllUser finds all User.
func (u UserRoleService) FindAllUser() ([]model.User, error) {
	users, err := u.UserRole.FindAllUser()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find users")
	}
	return users, nil
}
