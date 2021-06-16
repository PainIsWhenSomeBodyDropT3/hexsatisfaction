package repository

import (
	"database/sql"
	"strings"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
)

// AuthorRepo is a author repository.
type AuthorRepo struct {
	db *sql.DB
}

// NewAuthorRepo is a AuthorRepo constructor.
func NewAuthorRepo(db *sql.DB) *AuthorRepo {
	return &AuthorRepo{db: db}
}

// Create creates new author and returns id.
func (a AuthorRepo) Create(author model.Author) (int, error) {
	var creatID int
	rows, err := a.db.Query("INSERT INTO author (name, age, description, userID) VALUES ($1,$2,$3,$4) RETURNING id",
		author.Name, author.Age, author.Description, author.UserID)
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

// Update updates author and returns id.
func (a AuthorRepo) Update(id int, author model.Author) (int, error) {
	var updatedID int
	rows, err := a.db.Query("UPDATE author SET name=$1, age=$2, description=$3, userID=$4 WHERE id=$5 RETURNING id",
		author.Name, author.Age, author.Description, author.UserID, id)
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

// Delete deletes author and returns deleted id.
func (a AuthorRepo) Delete(id int) (int, error) {
	var delID int
	rows, err := a.db.Query("DELETE FROM author WHERE id=$1 RETURNING id ", id)
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

// FindByID finds author by id.
func (a AuthorRepo) FindByID(id int) (*model.Author, error) {
	var author model.Author
	rows, err := a.db.Query("SELECT * FROM author WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&author.ID, &author.Name, &author.Age, &author.Description, &author.UserID)
		if err != nil {
			return nil, err
		}
	}

	return &author, rows.Err()
}

// IsExistByID checks if author exist.
func (u AuthorRepo) IsExistByID(id int) (bool, error) {
	existingAuthor, err := u.FindByID(id)
	if err != nil && !strings.Contains(err.Error(), "sql: Rows are closed") {
		return false, err
	} else if existingAuthor.ID != 0 {
		return true, nil
	}

	return false, nil
}

// FindByUserID finds author by user id.
func (a AuthorRepo) FindByUserID(id int) (*model.Author, error) {
	var author model.Author
	rows, err := a.db.Query("SELECT * FROM author WHERE userID=$1", id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&author.ID, &author.Name, &author.Age, &author.Description, &author.UserID)
		if err != nil {
			return nil, err
		}
	}

	return &author, rows.Err()
}

// FindByName finds authors by name.
func (a AuthorRepo) FindByName(name string) ([]model.Author, error) {
	var authors []model.Author
	var author model.Author
	rows, err := a.db.Query("SELECT * FROM author WHERE name=$1", name)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&author.ID, &author.Name, &author.Age, &author.Description, &author.UserID)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, rows.Err()
}

// FindAll finds authors.
func (a AuthorRepo) FindAll() ([]model.Author, error) {
	var authors []model.Author
	var author model.Author
	rows, err := a.db.Query("SELECT * FROM author")
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&author.ID, &author.Name, &author.Age, &author.Description, &author.UserID)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, rows.Err()
}
