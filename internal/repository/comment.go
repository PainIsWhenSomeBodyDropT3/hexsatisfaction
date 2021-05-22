package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
)

// CommentRepo is a purchase repository.
type CommentRepo struct {
	db *sql.DB
}

// NewCommentRepo is a CommentRepo constructor.
func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

// Create creates comment and returns id.
func (c CommentRepo) Create(comment model.Comment) (int, error) {
	var creatID int
	rows, err := c.db.Query("INSERT INTO comment (userID, purchaseID, date, text) VALUES ($1,$2,$3,$4) RETURNING id ", comment.UserID, comment.PurchaseID, comment.Date, comment.Text)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&creatID)
		if err != nil {
			return 0, err
		}
	}
	return creatID, rows.Err()
}

// Update updates comment and returns id.
func (c CommentRepo) Update(id int, comment model.Comment) (int, error) {
	var updatedID int
	rows, err := c.db.Query("UPDATE comment SET userID=$1, purchaseID=$2, date=$3,text=$4 WHERE id = $5 RETURNING id", comment.UserID, comment.PurchaseID, comment.Date, comment.Text, id)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&updatedID)
		if err != nil {
			return 0, err
		}
	}
	return updatedID, rows.Err()

}

// Delete deletes comment and returns id.
func (c CommentRepo) Delete(id int) (int, error) {
	var delID int
	rows, err := c.db.Query("DELETE FROM comment WHERE id=$1 RETURNING id ", id)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&delID)
		if err != nil {
			return 0, err
		}
	}

	return delID, rows.Err()
}

// FindByID finds comment by id.
func (c CommentRepo) FindByID(id int) (*model.Comment, error) {
	var comment model.Comment
	rows, err := c.db.Query("SELECT * FROM comment WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.PurchaseID, &comment.Date, &comment.Text)
		if err != nil {
			return nil, err
		}
	}

	return &comment, rows.Err()
}

// FindAllByUserID finds comments by user id.
func (c CommentRepo) FindAllByUserID(id int) ([]model.Comment, error) {
	var comments []model.Comment
	var comment model.Comment
	rows, err := c.db.Query("SELECT * FROM comment WHERE userID=$1", id)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&comment.ID, &comment.UserID, &comment.PurchaseID, &comment.Date, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, rows.Err()
}

// FindByPurchaseID finds comments by purchase id.
func (c CommentRepo) FindByPurchaseID(id int) ([]model.Comment, error) {
	var comments []model.Comment
	var comment model.Comment
	rows, err := c.db.Query("SELECT * FROM comment WHERE purchaseID=$1", id)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&comment.ID, &comment.UserID, &comment.PurchaseID, &comment.Date, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, rows.Err()
}

// FindByUserIDAndPurchaseID finds comments by purchase and user id.
func (c CommentRepo) FindByUserIDAndPurchaseID(userID, purchaseID int) ([]model.Comment, error) {
	var comments []model.Comment
	var comment model.Comment
	rows, err := c.db.Query("SELECT * FROM comment WHERE userID=$1 AND purchaseID=$2", userID, purchaseID)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&comment.ID, &comment.UserID, &comment.PurchaseID, &comment.Date, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, rows.Err()
}

// FindAll finds comments.
func (c CommentRepo) FindAll() ([]model.Comment, error) {
	var comments []model.Comment
	var comment model.Comment
	rows, err := c.db.Query("SELECT * FROM comment")
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&comment.ID, &comment.UserID, &comment.PurchaseID, &comment.Date, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, rows.Err()
}

// FindByText finds comments by text.
func (c CommentRepo) FindByText(text string) ([]model.Comment, error) {
	var comments []model.Comment
	var comment model.Comment
	text = fmt.Sprintf("%%%s%%", text)
	rows, err := c.db.Query("SELECT * FROM comment WHERE text LIKE $1", text)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&comment.ID, &comment.UserID, &comment.PurchaseID, &comment.Date, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, rows.Err()
}

// FindByPeriod finds comments by date period.
func (c CommentRepo) FindByPeriod(start, end time.Time) ([]model.Comment, error) {
	var comments []model.Comment
	var comment model.Comment
	rows, err := c.db.Query("SELECT * FROM comment WHERE date BETWEEN $1 AND $2", start, end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&comment.ID, &comment.UserID, &comment.PurchaseID, &comment.Date, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, rows.Err()
}
