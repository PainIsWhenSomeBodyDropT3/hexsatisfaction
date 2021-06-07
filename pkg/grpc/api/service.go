package api

import (
	"context"

	"github.com/JesusG2000/hexsatisfaction/internal/repository"
)

type ExistChecker struct {
	repository.Repositories
}

func NewExistChecker(repositories repository.Repositories) *ExistChecker {
	return &ExistChecker{repositories}
}

func (e *ExistChecker) User(ctx context.Context, req *IsUserExistRequest) (*IsUserExistResponse, error) {
	result, err := e.Repositories.User.IsExistByID(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &IsUserExistResponse{
		Exist: result,
	}, nil
}

func (e *ExistChecker) Author(ctx context.Context, req *IsAuthorExistRequest) (*IsAuthorExistResponse, error) {
	result, err := e.Repositories.Author.IsExistByID(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &IsAuthorExistResponse{
		Exist: result,
	}, nil
}
