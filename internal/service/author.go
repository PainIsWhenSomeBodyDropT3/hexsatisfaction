package service

import (
	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/pkg/errors"
)

// AuthorService is a author service.
type AuthorService struct {
	repository.Author
}

// NewAuthorService is a AuthorService service constructor.
func NewAuthorService(author repository.Author) *AuthorService {
	return &AuthorService{author}
}

// Create creates author and returns id.
func (a AuthorService) Create(request model.CreateAuthorRequest) (int, error) {
	author := model.Author{
		Name:        request.Name,
		Age:         request.Age,
		Description: request.Description,
		UserID:      request.UserID,
	}
	id, err := a.Author.Create(author)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't create author")
	}

	return id, nil
}

// Update updates author and returns id.
func (a AuthorService) Update(request model.UpdateAuthorRequest) (int, error) {
	author := model.Author{
		Name:        request.Name,
		Age:         request.Age,
		Description: request.Description,
		UserID:      request.UserID,
	}
	id, err := a.Author.Update(request.ID, author)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't update author")
	}

	return id, nil
}

// Delete deletes author and returns deleted id.
func (a AuthorService) Delete(request model.DeleteAuthorRequest) (int, error) {
	id, err := a.Author.Delete(request.ID)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't delete author")
	}

	return id, nil
}

// FindByID finds author by id.
func (a AuthorService) FindByID(request model.IDAuthorRequest) (*model.Author, error) {
	author, err := a.Author.FindByID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find author")
	}

	return author, nil
}

// FindByUserID finds author by user id.
func (a AuthorService) FindByUserID(request model.UserIDAuthorRequest) (*model.Author, error) {
	author, err := a.Author.FindByUserID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find author")
	}

	return author, nil
}

// FindByName finds authors by name.
func (a AuthorService) FindByName(request model.NameAuthorRequest) ([]model.Author, error) {
	authors, err := a.Author.FindByName(request.Name)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find authors")
	}

	return authors, nil
}

// FindAll finds authors.
func (a AuthorService) FindAll() ([]model.Author, error) {
	authors, err := a.Author.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find authors")
	}

	return authors, nil
}
