package service

import (
	"testing"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	testAssert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_FindByLogin(t *testing.T) {
	assert := testAssert.New(t)
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
			fn: func(user *m.User, data test) {
				user.On("FindByLogin", data.login).
					Return(data.expRes, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find a user by login"),
		},
		{
			name:  "All ok",
			login: "test",
			fn: func(user *m.User, data test) {
				user.On("FindByLogin", data.login).
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
			user := new(m.User)
			service := NewUserService(user, api.TokenManager)
			if tc.fn != nil {
				tc.fn(user, tc)
			}
			u, err := service.FindByLogin(tc.login)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expRes, u)
		})
	}
}

func TestUser_FindByCredentials(t *testing.T) {
	assert := testAssert.New(t)
	api, err := InitTest4Mock()
	require.NoError(t, err)
	type test struct {
		name   string
		req    model.LoginUserRequest
		fn     func(user *m.User, data test)
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
			fn: func(user *m.User, data test) {
				user.On("FindByCredentials", model.User{
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
			fn: func(user *m.User, data test) {
				user.On("FindByCredentials", model.User{
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
			user := new(m.User)
			service := NewUserService(user, api.TokenManager)
			if tc.fn != nil {
				tc.fn(user, tc)
			}
			token, err := service.FindByCredentials(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			} else {
				assert.NotEmpty(token)
			}
		})
	}
}

func TestUser_IsExist(t *testing.T) {
	assert := testAssert.New(t)
	api, err := InitTest4Mock()
	require.NoError(t, err)
	type test struct {
		name   string
		login  string
		fn     func(user *m.User, data test)
		expRes bool
		expErr error
	}
	tt := []test{
		{
			name:  "IsExist errors",
			login: "test",
			fn: func(user *m.User, data test) {
				user.On("IsExist", data.login).
					Return(data.expRes, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't check user existence"),
		},
		{
			name:  "All ok",
			login: "test",
			fn: func(user *m.User, data test) {
				user.On("IsExist", data.login).
					Return(data.expRes, nil)
			},
			expRes: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			user := new(m.User)
			service := NewUserService(user, api.TokenManager)
			if tc.fn != nil {
				tc.fn(user, tc)
			}
			exist, err := service.IsExist(tc.login)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expRes, exist)
		})
	}
}

func TestUser_Create(t *testing.T) {
	assert := testAssert.New(t)
	api, err := InitTest4Mock()
	require.NoError(t, err)
	type test struct {
		name   string
		req    model.RegisterUserRequest
		fn     func(user *m.User, data test)
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
			fn: func(user *m.User, data test) {
				user.On("Create", model.User{
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
			fn: func(user *m.User, data test) {
				user.On("Create", model.User{
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
			user := new(m.User)
			service := NewUserService(user, api.TokenManager)
			if tc.fn != nil {
				tc.fn(user, tc)
			}
			id, err := service.Create(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}
