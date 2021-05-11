package repository

import (
	"database/sql"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
)

// UserRoleRepo is a user role repository.
type UserRoleRepo struct {
	db *sql.DB
}

// NewUserRoleRepo is a UserRoleRepo constructor.
func NewUserRoleRepo(db *sql.DB) *UserRoleRepo {
	return &UserRoleRepo{db: db}
}

// FindAllUser finds all User.
func (u UserRoleRepo) FindAllUser() ([]model.User, error) {
	var users []model.User
	var user model.User
	rows, err := u.db.Query("SELECT u.id , u.login , u.password , u.roleID FROM users u INNER JOIN user_role ur ON u.roleID=ur.id WHERE u.roleID=$1", dto.USER)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&user.ID, &user.Login, &user.Password, &user.RoleID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}
