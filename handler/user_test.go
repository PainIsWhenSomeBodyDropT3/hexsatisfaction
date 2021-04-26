package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/JesusG2000/hexsatisfaction/handler/mock"
	"github.com/JesusG2000/hexsatisfaction/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const slash = "/"

func TestUser_Login(t *testing.T) {
	id := 23
	tt := []struct {
		name     string
		path     string
		method   string
		isOkBody bool
		isOkRes  bool
		fn       func(userService *m.UserService)
		expCode  int
		expBody  model.User
	}{
		{
			name:    "bad body",
			path:    userPath + slash,
			method:  http.MethodPost,
			expCode: http.StatusBadRequest,
		},
		{
			name:     "no user",
			path:     userPath + slash,
			method:   http.MethodPost,
			isOkBody: true,
			fn: func(userService *m.UserService) {
				userService.On("FindByCredentials", mock.Anything).
					Return(&model.User{}, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:     "all ok",
			path:     userPath + slash,
			method:   http.MethodPost,
			isOkBody: true,
			isOkRes:  true,
			fn: func(userService *m.UserService) {
				userService.On("FindByCredentials", mock.Anything).
					Return(&model.User{
						ID:       id,
						Login:    "test",
						Password: "test",
					}, nil)
			},
			expCode: http.StatusOK,
			expBody: model.User{
				ID:       id,
				Login:    "test",
				Password: "test",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var user model.User
			var userLogin model.LoginUserRequest
			userService := new(m.UserService)
			service := newUser(userService)
			if tc.fn != nil {
				tc.fn(userService)
			}

			if tc.isOkBody {
				userLogin.Login = "test"
				userLogin.Password = "test"
			}

			payloadBuf := new(bytes.Buffer)
			err := json.NewEncoder(payloadBuf).Encode(&userLogin)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path, payloadBuf)
			require.NoError(t, err)

			res := httptest.NewRecorder()
			service.loginUser(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkBody && tc.isOkRes {
				err := json.NewDecoder(res.Body).Decode(&user)
				require.NoError(t, err)
			}
			require.Equal(t, tc.expBody, user)
		})
	}
}

func TestUser_Registration(t *testing.T) {
	tt := []struct {
		name     string
		path     string
		method   string
		isOkBody bool
		isOkRes  bool
		fn       func(userService *m.UserService)
		expCode  int
		expBody  string
	}{
		{
			name:    "bad body",
			path:    userPath + slash,
			method:  http.MethodPost,
			expCode: http.StatusBadRequest,
		},
		{
			name:     "existed user",
			path:     userPath + slash,
			method:   http.MethodPost,
			isOkBody: true,
			isOkRes:  true,
			fn: func(userService *m.UserService) {
				userService.On("IsExist", mock.Anything).
					Return(true, nil)
			},
			expCode: http.StatusFound,
			expBody: "this user already exist",
		},
		{
			name:     "all ok",
			path:     userPath + slash,
			method:   http.MethodPost,
			isOkBody: true,
			isOkRes:  true,
			fn: func(userService *m.UserService) {
				userService.On("IsExist", mock.Anything).
					Return(false, nil)
				userService.On("Create", mock.Anything).
					Return(nil)
			},
			expCode: http.StatusOK,
			expBody: "user created",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var userRes string
			var userLogin model.RegisterUserRequest
			userService := new(m.UserService)
			service := newUser(userService)
			if tc.fn != nil {
				tc.fn(userService)
			}

			if tc.isOkBody {
				userLogin.Login = "test"
				userLogin.Password = "test"
			}

			payloadBuf := new(bytes.Buffer)
			err := json.NewEncoder(payloadBuf).Encode(&userLogin)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path, payloadBuf)
			require.NoError(t, err)

			res := httptest.NewRecorder()
			service.registerUser(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkBody && tc.isOkRes {
				err := json.NewDecoder(res.Body).Decode(&userRes)
				require.NoError(t, err)
			}
			require.Equal(t, tc.expBody, userRes)
		})
	}
}
