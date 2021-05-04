package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	m "github.com/JesusG2000/hexsatisfaction/internal/handler/mock"
	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const slash = "/"
const login = "login"
const registration = "registration"

//NEED TO FIX

func TestUser_Login(t *testing.T) {
	api, err := service.InitTest4Mock()
	require.NoError(t, err)
	id := 15
	token, err := api.TokenManager.NewJWT(strconv.Itoa(id))
	tt := []struct {
		name     string
		path     string
		method   string
		isOkBody bool
		isOkRes  bool
		fn       func(userService *m.User)
		expCode  int
		expBody  string
	}{
		{
			name:    "bad body",
			path:    userPath + slash + login,
			method:  http.MethodPost,
			expCode: http.StatusBadRequest,
		},
		{
			name:     "no user",
			path:     userPath + slash + login,
			method:   http.MethodPost,
			isOkBody: true,
			fn: func(userService *m.User) {
				userService.On("FindByCredentials", mock.Anything).
					Return("", nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:     "all ok",
			path:     userPath + slash + login,
			method:   http.MethodPost,
			isOkBody: true,
			isOkRes:  true,
			fn: func(userService *m.User) {
				userService.On("FindByCredentials", mock.Anything).
					Return(token, nil)
			},
			expCode: http.StatusOK,
			expBody: token,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var token string
			var userLogin model.LoginUserRequest
			userService := new(m.User)
			router := newUser(api.Services, api.TokenManager)
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
			router.loginUser(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkBody && tc.isOkRes {
				err := json.NewDecoder(res.Body).Decode(&token)
				require.NoError(t, err)
			}
			require.Equal(t, tc.expBody, token)
		})
	}
}

func TestUser_Registration(t *testing.T) {
	api, err := service.InitTest4Mock()
	require.NoError(t, err)
	id := 23
	tt := []struct {
		name     string
		path     string
		method   string
		isOkBody bool
		isOkRes  bool
		fn       func(userService *m.User)
		expCode  int
		expBody  string
	}{
		{
			name:    "bad body",
			path:    userPath + slash + registration,
			method:  http.MethodPost,
			expCode: http.StatusBadRequest,
		},
		{
			name:     "existed user",
			path:     userPath + slash + registration,
			method:   http.MethodPost,
			isOkBody: true,
			isOkRes:  true,
			fn: func(userService *m.User) {
				userService.On("IsExist", mock.Anything).
					Return(true, nil)
			},
			expCode: http.StatusFound,
			expBody: "this user already exist",
		},
		{
			name:     "all ok",
			path:     userPath + slash + registration,
			method:   http.MethodPost,
			isOkBody: true,
			isOkRes:  true,
			fn: func(userService *m.User) {
				userService.On("IsExist", mock.Anything).
					Return(false, nil)
				userService.On("Create", mock.Anything).
					Return(id, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(id),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var userRes string
			var userLogin model.RegisterUserRequest
			userService := new(m.User)
			router := newUser(api.Services, api.TokenManager)
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
			router.registerUser(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkBody && tc.isOkRes {
				err := json.NewDecoder(res.Body).Decode(&userRes)
				require.NoError(t, err)
			}
			require.Equal(t, tc.expBody, userRes)
		})
	}
}
