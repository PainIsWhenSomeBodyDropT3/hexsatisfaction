package service

import (
	"testing"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_FindByLogin(t *testing.T) {
	a := assert.New(t)
	api, err := InitTest4Mock()
	require.NoError(t, err)
	type test struct {
		name   string
		login  string
		fn     func(user *m.User, data test)
		expRes *model.User
		expErr error
	}
	tt := []test{
		{
			name:  "FindByID errors",
			login: "test",
			fn: func(userDB *m.User, data test) {
				userDB.On("FindByLogin", data.login).
					Return(data.expRes, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find a user by login"),
		},
		{
			name:  "All ok",
			login: "test",
			fn: func(userDB *m.User, data test) {
				userDB.On("FindByLogin", data.login).
					Return(data.expRes, nil)
			},
			expRes: &model.User{
				ID:       15,
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.User)
			service := NewUserService(userDB, api.TokenManager)
			if tc.fn != nil {
				tc.fn(userDB, tc)
			}
			user, err := service.FindByLogin(tc.login)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expRes, user)
		})
	}
}

func TestUser_FindByCredentials(t *testing.T) {
	a := assert.New(t)
	api, err := InitTest4Mock()
	require.NoError(t, err)
	type test struct {
		name   string
		req    model.LoginUserRequest
		fn     func(userDB *m.User, data test)
		expRes *model.User
		expErr error
	}
	tt := []test{
		{
			name: "FindByCredentials errors",
			req: model.LoginUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userDB *m.User, data test) {
				userDB.On("FindByCredentials", model.User{
					Login:    data.req.Login,
					Password: data.req.Password,
				}).Return(data.expRes, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find a user by credentials"),
		},
		{
			name: "All ok",
			req: model.LoginUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userDB *m.User, data test) {
				userDB.On("FindByCredentials", model.User{
					Login:    data.req.Login,
					Password: data.req.Password,
				}).Return(data.expRes, nil)
			},
			expRes: &model.User{
				ID:       15,
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.User)
			service := NewUserService(userDB, api.TokenManager)
			if tc.fn != nil {
				tc.fn(userDB, tc)
			}
			token, err := service.FindByCredentials(tc.req)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			} else {
				a.NotEmpty(token)
			}
		})
	}
}

func TestUser_IsExist(t *testing.T) {
	a := assert.New(t)
	api, err := InitTest4Mock()
	require.NoError(t, err)
	type test struct {
		name   string
		login  string
		fn     func(userDB *m.User, data test)
		expRes bool
		expErr error
	}
	tt := []test{
		{
			name:  "IsExist errors",
			login: "test",
			fn: func(userDB *m.User, data test) {
				userDB.On("IsExist", data.login).
					Return(data.expRes, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't check user existence"),
		},
		{
			name:  "All ok",
			login: "test",
			fn: func(userDB *m.User, data test) {
				userDB.On("IsExist", data.login).
					Return(data.expRes, nil)
			},
			expRes: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.User)
			service := NewUserService(userDB, api.TokenManager)
			if tc.fn != nil {
				tc.fn(userDB, tc)
			}
			user, err := service.IsExist(tc.login)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expRes, user)
		})
	}
}

func TestUser_Create(t *testing.T) {
	a := assert.New(t)
	api, err := InitTest4Mock()
	require.NoError(t, err)
	type test struct {
		name   string
		req    model.RegisterUserRequest
		fn     func(userDB *m.User, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Create errors",
			req: model.RegisterUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userDB *m.User, data test) {
				userDB.On("Create", model.User{
					Login:    data.req.Login,
					Password: data.req.Password,
				}).
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
			fn: func(userDB *m.User, data test) {
				userDB.On("Create", model.User{
					Login:    data.req.Login,
					Password: data.req.Password,
				}).
					Return(data.expID, nil)
			},
			expID: 15,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userDB := new(m.User)
			service := NewUserService(userDB, api.TokenManager)
			if tc.fn != nil {
				tc.fn(userDB, tc)
			}
			id, err := service.Create(tc.req)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expID, id)
		})
	}
}
