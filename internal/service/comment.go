package service

import (
	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/pkg/errors"
)

// CommentService is a purchase service.
type CommentService struct {
	repository.Comment
}

// NewCommentService is a CommentService service constructor.
func NewCommentService(comment repository.Comment) *CommentService {
	return &CommentService{comment}
}

// Create creates comments and returns id.
func (c CommentService) Create(request model.CreateCommentRequest) (int, error) {
	comment := model.Comment{
		UserID:     request.UserID,
		PurchaseID: request.PurchaseID,
		Date:       request.Date,
		Text:       request.Text,
	}
	id, err := c.Comment.Create(comment)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't create comment")
	}

	return id, nil
}

// Update updates comments and returns id.
func (c CommentService) Update(request model.UpdateCommentRequest) (int, error) {
	comment := model.Comment{
		UserID:     request.UserID,
		PurchaseID: request.PurchaseID,
		Date:       request.Date,
		Text:       request.Text,
	}
	id, err := c.Comment.Update(request.ID, comment)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't update comment")
	}

	return id, nil
}

// Delete deletes comments and returns id.
func (c CommentService) Delete(request model.DeleteCommentRequest) (int, error) {
	id, err := c.Comment.Delete(request.ID)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't delete comment")
	}

	return id, nil
}

// FindByID finds comments by id.
func (c CommentService) FindByID(request model.IDCommentRequest) (*model.Comment, error) {
	comment, err := c.Comment.FindByID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find comment")
	}

	return comment, nil
}

// FindAllByUserID finds comments by user id.
func (c CommentService) FindAllByUserID(request model.UserIDCommentRequest) ([]model.Comment, error) {
	comments, err := c.Comment.FindAllByUserID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find comments")
	}

	return comments, nil
}

// FindByPurchaseID finds comments by purchase id.
func (c CommentService) FindByPurchaseID(request model.PurchaseIDCommentRequest) ([]model.Comment, error) {
	comments, err := c.Comment.FindByPurchaseID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find comments")
	}

	return comments, nil
}

// FindByUserIDAndPurchaseID finds comments by purchase and user id.
func (c CommentService) FindByUserIDAndPurchaseID(request model.UserPurchaseIDCommentRequest) ([]model.Comment, error) {
	comments, err := c.Comment.FindByUserIDAndPurchaseID(request.UserID, request.PurchaseID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find comments")
	}

	return comments, nil
}

// FindAll finds all comments.
func (c CommentService) FindAll() ([]model.Comment, error) {
	comments, err := c.Comment.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find comments")
	}

	return comments, nil
}

// FindByText finds comments by text.
func (c CommentService) FindByText(request model.TextCommentRequest) ([]model.Comment, error) {
	comments, err := c.Comment.FindByText(request.Text)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find comments")
	}

	return comments, nil
}

// FindByPeriod finds comments by date period.
func (c CommentService) FindByPeriod(request model.PeriodCommentRequest) ([]model.Comment, error) {
	comments, err := c.Comment.FindByPeriod(request.Start, request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find comments")
	}

	return comments, nil
}
