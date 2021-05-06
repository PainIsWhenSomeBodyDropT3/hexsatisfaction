package repository

import (
	"testing"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	"github.com/stretchr/testify/require"
)

func TestUserRole_FindAllUser(t *testing.T) {
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		isOk     bool
		expUsers []model.User
	}{
		{
			name:     "user not found errors",
			expUsers: nil,
		},
		{
			name: "all ok",
			isOk: true,
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
			require.NoError(t, err)
			if tc.isOk {
				id1, err := repos.User.Create(model.User{
					Login:    "test",
					Password: "test",
				})
				require.NoError(t, err)
				id2, err := repos.User.Create(model.User{
					Login:    "test1",
					Password: "test1",
				})
				require.NoError(t, err)

				tc.expUsers[0].ID = id1
				tc.expUsers[1].ID = id2

			}
			users, err := repos.UserRole.FindAllUser()
			require.NoError(t, err)
			require.Equal(t, tc.expUsers, users)
			if tc.isOk {
				_, err := db.Exec("DELETE FROM users")
				require.NoError(t, err)
			}
		})
	}
}
