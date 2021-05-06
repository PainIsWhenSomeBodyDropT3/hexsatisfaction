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
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	slash        = "/"
	login        = "login"
	registration = "registration"
	api          = "api"
	getAll       = "getAll"
)

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
			api.Services.User = userService
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
		}, {
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
			api.Services.User = userService
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

func TestUserRole_FindAll(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	id := 15
	_, err = testApi.TokenManager.NewJWT(strconv.Itoa(id))
	tt := []struct {
		name    string
		path    string
		method  string
		fn      func(userRoleService *m.UserRole)
		expCode int
		expBody []model.User
	}{
		{
			name:   "all ok",
			path:   userPath + slash + api + slash + getAll,
			method: http.MethodGet,
			fn: func(userRoleService *m.UserRole) {
				userRoleService.On("FindAllUser").
					Return([]model.User{
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
					}, nil)
			},
			expCode: http.StatusOK,
			expBody: []model.User{
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
			var users []model.User
			userRoleService := new(m.UserRole)
			testApi.Services.UserRole = userRoleService
			router := newUser(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(userRoleService)
			}

			req, err := http.NewRequest(tc.method, tc.path, nil)
			require.NoError(t, err)

			res := httptest.NewRecorder()
			router.getAllUser(res, req)
			require.Equal(t, tc.expCode, res.Code)

			err = json.NewDecoder(res.Body).Decode(&users)
			require.NoError(t, err)

			require.Equal(t, tc.expBody, users)
		})
	}
}
