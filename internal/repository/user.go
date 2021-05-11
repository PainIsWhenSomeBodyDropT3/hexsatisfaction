package repository

import (
	"database/sql"
	"strings"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
)

// UserRepo is a user repository.
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo is a UserRepo constructor.
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create saves user and returns id.
func (u UserRepo) Create(user model.User) (int, error) {
	var id int
	rows, err := u.db.Query("INSERT INTO users (login , password,roleID) VALUES ($1,$2,$3) RETURNING id ", user.Login, user.Password, dto.USER)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, rows.Err()
}

// FindByLogin finds the User by login.
func (u UserRepo) FindByLogin(login string) (*model.User, error) {
	var user model.User
	rows, err := u.db.Query("SELECT * FROM users WHERE login = $1", login)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Login, &user.Password, &user.RoleID)
		if err != nil {
			return nil, err
		}
	}

	return &user, rows.Err()
}

// FindByCredentials finds the User by credentials.
func (u UserRepo) FindByCredentials(user model.User) (*model.User, error) {
	var newUser model.User
	rows, err := u.db.Query("SELECT * FROM users WHERE login = $1 AND password = $2", user.Login, user.Password)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&newUser.ID, &newUser.Login, &newUser.Password, &newUser.RoleID)
		if err != nil {
			return nil, err
		}
	}

	return &newUser, rows.Err()
}

// IsExist checks if User exist.
func (u UserRepo) IsExist(login string) (bool, error) {
	existingUser, err := u.FindByLogin(login)
	if err != nil && !strings.Contains(err.Error(), "sql: Rows are closed") {
		return false, err
	} else if existingUser.ID != 0 {
		return true, nil
	}

	return false, nil
}
