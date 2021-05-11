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
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	user         = "user"
	slash        = "/"
	login        = "login"
	registration = "registration"
	api          = "api"
	getAll       = "getAll"
)
const authorizationHeader = "Authorization"

func TestUser_Login(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)

	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)
	type test struct {
		name     string
		path     string
		method   string
		req      model.LoginUserRequest
		isNoBody bool
		fn       func(userService *m.User, data test)
		expCode  int
		expBody  string
	}
	tt := []test{
		{
			name:   "invalid login",
			path:   slash + user + slash + login,
			method: http.MethodPost,
			req: model.LoginUserRequest{
				Login:    "",
				Password: "test",
			},
			fn: func(userService *m.User, data test) {
				userService.On("FindByCredentials", data.req).
					Return(data.expBody, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "login is required",
		},
		{
			name:   "find err",
			path:   slash + user + slash + login,
			method: http.MethodPost,
			req: model.LoginUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userService *m.User, data test) {
				userService.On("FindByCredentials", data.req).
					Return(data.expBody, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "no user",
			path:   slash + user + slash + login,
			method: http.MethodPost,
			req: model.LoginUserRequest{
				Login:    "test",
				Password: "test",
			},
			isNoBody: true,
			fn: func(userService *m.User, data test) {
				userService.On("FindByCredentials", data.req).
					Return(data.expBody, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:   "all ok",
			path:   slash + user + slash + login,
			method: http.MethodPost,
			req: model.LoginUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userService *m.User, data test) {
				userService.On("FindByCredentials", data.req).
					Return(data.expBody, nil)
			},
			expCode: http.StatusOK,
			expBody: token,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			userService := new(m.User)
			testApi.Services.User = userService
			router := newUser(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(userService, tc)
			}

			payloadBuf := new(bytes.Buffer)
			err := json.NewEncoder(payloadBuf).Encode(&tc.req)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path, payloadBuf)
			require.NoError(t, err)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if !tc.isNoBody {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
			}
			require.Equal(t, tc.expBody, r)
		})
	}
}

func TestUser_Registration(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	type test struct {
		name    string
		path    string
		method  string
		req     model.RegisterUserRequest
		fn      func(userService *m.User, data test)
		expCode int
		expBody string
	}
	tt := []test{
		{
			name:   "bad login",
			path:   slash + user + slash + registration,
			method: http.MethodPost,
			req: model.RegisterUserRequest{
				Login:    "",
				Password: "test",
			},
			fn: func(userService *m.User, data test) {
				userService.On("IsExist", data.req.Login).
					Return(false, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "login is required",
		},
		{
			name:   "exist error",
			path:   slash + user + slash + registration,
			method: http.MethodPost,
			req: model.RegisterUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userService *m.User, data test) {
				userService.On("IsExist", data.req.Login).
					Return(false, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "existed user",
			path:   slash + user + slash + registration,
			method: http.MethodPost,
			req: model.RegisterUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userService *m.User, data test) {
				userService.On("IsExist", data.req.Login).
					Return(true, nil)
			},
			expCode: http.StatusFound,
			expBody: "this user already exist",
		},
		{
			name:   "create error",
			path:   slash + user + slash + registration,
			method: http.MethodPost,
			req: model.RegisterUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userService *m.User, data test) {
				userService.On("IsExist", data.req.Login).
					Return(false, nil)
				userService.On("Create", data.req).
					Return(0, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "all ok",
			path:   slash + user + slash + registration,
			method: http.MethodPost,
			req: model.RegisterUserRequest{
				Login:    "test",
				Password: "test",
			},
			fn: func(userService *m.User, data test) {
				userService.On("IsExist", data.req.Login).
					Return(false, nil)
				userService.On("Create", data.req).
					Return(15, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(15),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			userService := new(m.User)
			testApi.Services.User = userService
			router := newUser(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(userService, tc)
			}

			payloadBuf := new(bytes.Buffer)
			err := json.NewEncoder(payloadBuf).Encode(&tc.req)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path, payloadBuf)
			require.NoError(t, err)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			err = json.NewDecoder(res.Body).Decode(&r)
			require.NoError(t, err)
			require.Equal(t, tc.expBody, r)
		})
	}
}

func TestUserRole_FindAll(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)

	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name    string
		path    string
		method  string
		isOkRes bool
		fn      func(userRoleService *m.UserRole, data test)
		expCode int
		expBody []model.User
	}
	tt := []test{

		{
			name:   "find error",
			path:   slash + user + slash + api + slash + getAll,
			method: http.MethodGet,
			fn: func(userRoleService *m.UserRole, data test) {
				userRoleService.On("FindAllUser").
					Return(data.expBody, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:    "all ok",
			path:    slash + user + slash + api + slash + getAll,
			method:  http.MethodGet,
			isOkRes: true,
			fn: func(userRoleService *m.UserRole, data test) {
				userRoleService.On("FindAllUser").
					Return(data.expBody, nil)
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
			var r []model.User
			userRoleService := new(m.UserRole)
			testApi.Services.UserRole = userRoleService
			router := newUser(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(userRoleService, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path, nil)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)
			if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
			}

			require.Equal(t, tc.expBody, r)
		})
	}
}
