package repository

import (
	"database/sql"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
)

// User is an interface for UserRepo methods.
type User interface {
	Create(user model.User) (int, error)
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(user model.User) (*model.User, error)
	IsExist(login string) (bool, error)
}

// UserRole is an interface for UserRoleRepo methods.
type UserRole interface {
	FindAllUser() ([]model.User, error)
}

// Purchase is an interface for PurchaseRepo methods.
type Purchase interface {
	Create(purchase model.Purchase) (int, error)
	Delete(id int) (int, error)
	FindByID(id int) (*model.Purchase, error)
	FindLastByUserID(id int) (*model.Purchase, error)
	FindAllByUserID(id int) ([]model.Purchase, error)
	FindByUserIDAndPeriod(id int, start, end time.Time) ([]model.Purchase, error)
	FindByUserIDAfterDate(id int, start time.Time) ([]model.Purchase, error)
	FindByUserIDBeforeDate(id int, end time.Time) ([]model.Purchase, error)
	FindByUserIDAndFileID(userID, fileID int) ([]model.Purchase, error)
	FindLast() (*model.Purchase, error)
	FindAll() ([]model.Purchase, error)
	FindByPeriod(start, end time.Time) ([]model.Purchase, error)
	FindAfterDate(start time.Time) ([]model.Purchase, error)
	FindBeforeDate(end time.Time) ([]model.Purchase, error)
	FindByFileID(id int) ([]model.Purchase, error)
}

// Comment is an interface for CommentRepo methods.
type Comment interface {
	Create(comment model.Comment) (int, error)
	Update(id int, comment model.Comment) (int, error)
	Delete(id int) (int, error)
	DeleteByPurchaseID(id int) (int, error)
	FindByID(id int) (*model.Comment, error)
	FindAllByUserID(id int) ([]model.Comment, error)
	FindByPurchaseID(id int) ([]model.Comment, error)
	FindByUserIDAndPurchaseID(userID, purchaseID int) ([]model.Comment, error)
	FindAll() ([]model.Comment, error)
	FindByText(text string) ([]model.Comment, error)
	FindByPeriod(start, end time.Time) ([]model.Comment, error)
}

// File is an interface for FileRepo methods.
type File interface {
	Create(file model.File) (int, error)
	Update(id int, file model.File) (int, error)
	Delete(id int) (int, error)
	FindByID(id int) (*model.File, error)
	FindByName(name string) ([]model.File, error)
	FindAll() ([]model.File, error)
	FindByAuthorID(id int) ([]model.File, error)
	FindNotActual() ([]model.File, error)
	FindActual() ([]model.File, error)
	FindAddedByPeriod(start, end time.Time) ([]model.File, error)
	FindUpdatedByPeriod(start, end time.Time) ([]model.File, error)
}

// Author is an interface for AuthorRepo methods.
type Author interface {
	Create(author model.Author) (int, error)
	Update(id int, author model.Author) (int, error)
	Delete(id int) (int, error)
	FindByID(id int) (*model.Author, error)
	FindByUserID(id int) (*model.Author, error)
	FindByName(name string) ([]model.Author, error)
	FindAll() ([]model.Author, error)
}

// Repositories collects all repository interfaces.
type Repositories struct {
	User     User
	UserRole UserRole
	Purchase Purchase
	Comment  Comment
	File     File
	Author   Author
}

// NewRepositories is a Repositories constructor.
func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User:     NewUserRepo(db),
		UserRole: NewUserRoleRepo(db),
		Purchase: NewPurchaseRepo(db),
		Comment:  NewCommentRepo(db),
		File:     NewFileRepo(db),
		Author:   NewAuthorRepo(db),
	}
}
