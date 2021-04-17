package pg

import (
	"strings"

	"github.com/JesusG2000/hexsatisfaction/model"
)

// Create saves user.
func (u User) Create(user model.User) error {
	_, err := u.db.Exec("INSERT INTO users (login , password) VALUES ($1,$2)", user.Login, user.Password)

	return err
}

// FindByLogin find User by login.
func (u User) FindByLogin(login string) (*model.User, error) {
	var user model.User
	rows, err := u.db.Query("SELECT * FROM users WHERE login = $1", login)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			return nil, err
		}
	}

	return &user, rows.Err()
}

// FindByCredentials find User by credentials.
func (u User) FindByCredentials(user model.User) (*model.User, error) {
	var newUser model.User
	rows, err := u.db.Query("SELECT * FROM users WHERE login = $1 AND password = $2", user.Login, user.Password)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&newUser.ID, &newUser.Login, &newUser.Password)
		if err != nil {
			return nil, err
		}
	}

	return &newUser, rows.Err()
}

// IsExist checks is User exist.
func (u User) IsExist(login string) (bool, error) {
	existingUser, err := u.FindByLogin(login)
	if err != nil && !strings.Contains(err.Error(), "sql: Rows are closed") {
		return false, err
	} else if existingUser != nil {
		return true, nil
	}

	return false, nil
}
