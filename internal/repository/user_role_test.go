package repository

import (
	"testing"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRole_FindAllUser(t *testing.T) {
	a := assert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		isOk     bool
		users    []model.User
		expUsers []model.User
	}{
		{
			name: "user not found errors",
		},
		{
			name: "all ok",
			isOk: true,
			users: []model.User{
				{
					Login:    "test",
					Password: "test",
					RoleID:   dto.USER,
				},
				{
					Login:    "test1",
					Password: "test1",
					RoleID:   dto.USER,
				},
			},
			expUsers: []model.User{
				{
					Login:    "test",
					Password: "test",
					RoleID:   dto.USER,
				},
				{
					Login:    "test1",
					Password: "test1",
					RoleID:   dto.USER,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM users")
			a.Nil(err)
			if tc.isOk {
				for i := range tc.users {
					id, err := repos.User.Create(tc.users[i])
					a.Nil(err)
					tc.expUsers[i].ID = id
				}
			}
			users, err := repos.UserRole.FindAllUser()
			a.Nil(err)
			a.Equal(tc.expUsers, users)
			if tc.isOk {
				_, err := db.Exec("DELETE FROM users")
				a.Nil(err)
			}
		})
	}
	err = db.Close()
	require.NoError(t, err)
}
