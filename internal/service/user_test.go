package service

import (
	"testing"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUser_FindByLogin(t *testing.T) {
	api, err := InitTest4Mock()
	require.NoError(t, err)
	id := 15
	login := "test"
	tt := []struct {
		name   string
		login  string
		fn     func(user *m.User)
		expRes *model.User
		expErr error
	}{
		{
			name:  "FindById errors",
			login: login,
			fn: func(userDB *m.User) {
				userDB.On("FindByLogin", login).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find a user by login"),
		},
		{
			name:  "All ok",
			login: login,
			fn: func(userDB *m.User) {
				userDB.On("FindByLogin", login).
					Return(&model.User{
						ID:       id,
						Login:    login,
						Password: login,
						RoleID:   dto.USER,
					}, nil)
			},
			expRes: &model.User{
				ID:       id,
				Login:    login,
				Password: login,
				RoleID:   dto.USER,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.User)
			service := NewUserService(userDB, api.TokenManager)
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
	api, err := InitTest4Mock()
	require.NoError(t, err)
	id := 23
	tt := []struct {
		name   string
		user   model.LoginUserRequest
		fn     func(userDB *m.User)
		expErr error
	}{
		{
			name: "FindByCredentials errors",
			fn: func(userDB *m.User) {
				userDB.On("FindByCredentials", mock.Anything).
					Return(&model.User{}, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find a user by credentials"),
		},
		{
			name: "All ok",
			user: model.LoginUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userDB *m.User) {
				userDB.On("FindByCredentials", model.User{
					Login:    "test",
					Password: "test",
				}).
					Return(&model.User{
						ID:       id,
						Login:    "test",
						Password: "test",
						RoleID:   dto.USER,
					}, nil)

			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.User)
			service := NewUserService(userDB, api.TokenManager)
			if tc.fn != nil {
				tc.fn(userDB)
			}
			token, err := service.FindByCredentials(tc.user)
			if err != nil {

				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.Nil(t, err)
				require.NotEmpty(t, token)
			}
		})
	}
}

func TestUser_IsExist(t *testing.T) {
	api, err := InitTest4Mock()
	require.NoError(t, err)
	login := "test"
	tt := []struct {
		name   string
		login  string
		fn     func(userDB *m.User)
		expRes bool
		expErr error
	}{

		{
			name:  "IsExist errors",
			login: login,
			fn: func(userDB *m.User) {
				userDB.On("IsExist", login).
					Return(false, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't check user existence"),
		},
		{
			name:  "All ok",
			login: login,
			fn: func(userDB *m.User) {
				userDB.On("IsExist", login).
					Return(true, nil)
			},
			expRes: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.User)
			service := NewUserService(userDB, api.TokenManager)
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
	api, err := InitTest4Mock()
	require.NoError(t, err)
	id := 23
	tt := []struct {
		name   string
		req    model.RegisterUserRequest
		fn     func(userDB *m.User)
		expId  int
		expErr error
	}{
		{
			name: "Create errors",
			fn: func(userDB *m.User) {
				userDB.On("Create", mock.Anything).
					Return(0, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't create a user"),
		},
		{
			name: "All ok",
			req: model.RegisterUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userDB *m.User) {
				userDB.On("Create", model.User{
					Login:    "test",
					Password: "test",
				}).
					Return(id, nil)
			},
			expId: id,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.User)
			service := NewUserService(userDB, api.TokenManager)
			if tc.fn != nil {
				tc.fn(userDB)
			}
			id, err := service.Create(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expId, id)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
