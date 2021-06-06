package service

import (
	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
)

// User is an interface for UserService methods.
type User interface {
	Create(req model.RegisterUserRequest) (int, error)
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(req model.LoginUserRequest) (string, error)
	IsExist(login string) (bool, error)
}

// UserRole is an interface for UserRoleService methods.
type UserRole interface {
	FindAllUser() ([]model.User, error)
}

// Author is an interface for AuthorService repository methods.
type Author interface {
	Create(request model.CreateAuthorRequest) (int, error)
	Update(request model.UpdateAuthorRequest) (int, error)
	Delete(request model.DeleteAuthorRequest) (int, error)
	FindByID(request model.IDAuthorRequest) (*model.Author, error)
	FindByUserID(request model.UserIDAuthorRequest) (*model.Author, error)
	FindByName(request model.NameAuthorRequest) ([]model.Author, error)
	FindAll() ([]model.Author, error)
}

// Services collects all service interfaces.
type Services struct {
	User     User
	UserRole UserRole
	Author   Author
}

// Deps represents dependencies for services.
type Deps struct {
	Repos        *repository.Repositories
	TokenManager auth.TokenManager
}

// NewServices is a Services constructor.
func NewServices(deps Deps) *Services {
	return &Services{
		User:     NewUserService(deps.Repos.User, deps.TokenManager),
		UserRole: NewUserRoleService(deps.Repos.UserRole),
		Author:   NewAuthorService(deps.Repos.Author),
	}
}
