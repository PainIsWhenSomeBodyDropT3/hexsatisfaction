package controller

import (
	"testing"

	m "github.com/JesusG2000/hexsatisfaction/controller/mock"
	"github.com/JesusG2000/hexsatisfaction/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUser_FindByLogin(t *testing.T) {
	id := 15
	login := "test"
	tt := []struct {
		name   string
		login  string
		fn     func(userDB *m.UserDB)
		expRes *model.User
		expErr error
	}{
		{
			name:  "FindById errors",
			login: login,
			fn: func(userDB *m.UserDB) {
				userDB.On("FindByLogin", login).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find a user by login"),
		},
		{
			name:  "All ok",
			login: login,
			fn: func(userDB *m.UserDB) {
				userDB.On("FindByLogin", login).
					Return(&model.User{
						ID:       id,
						Login:    login,
						Password: login,
					}, nil)
			},
			expRes: &model.User{
				ID:       id,
				Login:    login,
				Password: login,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.UserDB)
			service := NewUser(userDB)
			if tc.fn != nil {
				tc.fn(userDB)
			}
			user, err := service.FindByLogin(tc.login)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.Nil(t, err)
			}
			require.Equal(t, tc.expRes, user)
		})
	}
}

func TestUser_FindByCredentials(t *testing.T) {
	id := 15
	tt := []struct {
		name   string
		user   model.LoginUserRequest
		fn     func(userDB *m.UserDB)
		expRes *model.User
		expErr error
	}{
		{
			name: "FindByCredentials errors",
			fn: func(userDB *m.UserDB) {
				userDB.On("FindByCredentials", mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find a user by credentials"),
		},
		{
			name: "All ok",
			user: model.LoginUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userDB *m.UserDB) {
				userDB.On("FindByCredentials", model.User{
					Login:    "test",
					Password: "test",
				}).
					Return(&model.User{
						ID:       id,
						Login:    "test",
						Password: "test",
					}, nil)
			},
			expRes: &model.User{
				ID:       id,
				Login:    "test",
				Password: "test",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.UserDB)
			service := NewUser(userDB)
			if tc.fn != nil {
				tc.fn(userDB)
			}
			user, err := service.FindByCredentials(tc.user)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.Nil(t, err)
			}
			require.Equal(t, tc.expRes, user)
		})
	}
}

func TestUser_IsExist(t *testing.T) {
	login := "test"
	tt := []struct {
		name   string
		login  string
		fn     func(userDB *m.UserDB)
		expRes bool
		expErr error
	}{

		{
			name:  "IsExist errors",
			login: login,
			fn: func(userDB *m.UserDB) {
				userDB.On("IsExist", login).
					Return(false, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't check user existence"),
		},
		{
			name:  "All ok",
			login: login,
			fn: func(userDB *m.UserDB) {
				userDB.On("IsExist", login).
					Return(true, nil)
			},
			expRes: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.UserDB)
			service := NewUser(userDB)
			if tc.fn != nil {
				tc.fn(userDB)
			}
			user, err := service.IsExist(tc.login)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.Nil(t, err)
			}
			require.Equal(t, tc.expRes, user)
		})
	}
}

func TestUser_Create(t *testing.T) {
	tt := []struct {
		name   string
		req    model.RegisterUserRequest
		fn     func(userDB *m.UserDB)
		expErr error
	}{
		{
			name: "Create errors",
			fn: func(userDB *m.UserDB) {
				userDB.On("Create", mock.Anything).
					Return(errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't create a user"),
		},
		{
			name: "All ok",
			req: model.RegisterUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userDB *m.UserDB) {
				userDB.On("Create", model.User{
					Login:    "test",
					Password: "test",
				}).
					Return(nil)
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.UserDB)
			service := NewUser(userDB)
			if tc.fn != nil {
				tc.fn(userDB)
			}
			err := service.Create(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
