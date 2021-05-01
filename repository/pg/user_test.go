package pg

import (
	"database/sql"
	"testing"

	"github.com/JesusG2000/hexsatisfaction/model"
	"github.com/JesusG2000/hexsatisfaction/model/dto"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func connect2UserTestRepository() (*sql.DB, *User, error) {

	f, err := NewPgRepository()
	if err != nil {
		return nil, nil, err
	}
	return f.DB, f.NewUserRepository(), nil
}

func TestUser_FindByCredentials(t *testing.T) {
	db, userRepo, err := connect2UserTestRepository()
	require.NoError(t, err)
	tt := []struct {
		name    string
		isOk    bool
		user    model.User
		expUser *model.User
	}{
		{
			name: "user not found errors",
			user: model.User{
				Login:    "not correct",
				Password: "not correct",
			},
			expUser: &model.User{},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
			},
			expUser: &model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM users")
			require.NoError(t, err)
			if tc.isOk {
				err := userRepo.Create(tc.user)
				require.NoError(t, err)
			}
			user, err := userRepo.FindByCredentials(tc.user)
			require.NoError(t, err)
			user.ID = 0
			require.Equal(t, tc.expUser, user)
			if tc.isOk {
				_, err := db.Exec("DELETE FROM users")
				require.NoError(t, err)
			}
		})
	}
}

func TestUser_IsExist(t *testing.T) {
	db, userRepo, err := connect2UserTestRepository()
	require.NoError(t, err)
	user := model.User{
		Login:    "test",
		Password: "test",
	}
	tt := []struct {
		name   string
		isOk   bool
		login  string
		expRes bool
	}{
		{
			name:  "user not found errors",
			login: "not correct",
		},
		{
			name:   "all ok",
			login:  "test",
			isOk:   true,
			expRes: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM users")
			require.NoError(t, err)
			if tc.isOk {
				err := userRepo.Create(user)
				require.NoError(t, err)
			}
			exist, err := userRepo.IsExist(tc.login)
			require.NoError(t, err)
			require.Equal(t, tc.expRes, exist)
			if tc.isOk {
				_, err := db.Exec("DELETE FROM users")
				require.NoError(t, err)
			}
		})
	}
}

func TestUser_FindByLogin(t *testing.T) {
	db, userRepo, err := connect2UserTestRepository()
	require.NoError(t, err)
	tt := []struct {
		name  string
		isOk  bool
		login string
		user  *model.User
	}{
		{
			name:  "user not found errors",
			login: "not correct",
			user:  &model.User{},
		},
		{
			name:  "all ok",
			isOk:  true,
			login: "test",
			user: &model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM users")
			require.NoError(t, err)
			if tc.isOk {
				err := userRepo.Create(*tc.user)
				require.NoError(t, err)
			}
			user, err := userRepo.FindByLogin(tc.login)
			require.NoError(t, err)
			user.ID = 0
			require.Equal(t, tc.user, user)
			if tc.isOk {
				_, err := db.Exec("DELETE FROM users")
				require.NoError(t, err)
			}
		})
	}
}
