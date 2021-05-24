package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	m "github.com/JesusG2000/hexsatisfaction/internal/handler/mock"
	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/pkg/errors"
	testAssert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	actual  = "actual"
	expired = "expired"
	added   = "added"
	updated = "updated"
)

func TestFile_Create(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name    string
		path    string
		method  string
		req     model.CreateFileRequest
		fn      func(fileService *m.File, data test)
		expCode int
		expBody string
	}

	tt := []test{
		{
			name:   "invalid author id",
			path:   slash + file + slash + api + slash,
			method: http.MethodPost,
			req: model.CreateFileRequest{
				Name:        "some",
				Description: "some",
				Size:        1,
				Path:        "some",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				Actual:      true,
				AuthorID:    0,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Create", data.req).
					Return(0, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "not correct author id",
		},
		{
			name:   "create err",
			path:   slash + file + slash + api + slash,
			method: http.MethodPost,
			req: model.CreateFileRequest{
				Name:        "some",
				Description: "some",
				Size:        1,
				Path:        "some",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				Actual:      true,
				AuthorID:    1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Create", data.req).
					Return(0, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "all ok",
			path:   slash + file + slash + api + slash,
			method: http.MethodPost,
			req: model.CreateFileRequest{
				Name:        "some",
				Description: "some",
				Size:        1,
				Path:        "some",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				Actual:      true,
				AuthorID:    1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Create", data.req).
					Return(15, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(15),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			assert.Nil(err)

			req, err := http.NewRequest(tc.method, tc.path, body)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)

			err = json.NewDecoder(res.Body).Decode(&r)
			assert.Nil(err)
			assert.Equal(tc.expBody, r)
		})
	}
}

func TestFile_Update(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name    string
		path    string
		method  string
		isOkRes bool
		req     model.UpdateFileRequest
		fn      func(fileService *m.File, data test)
		expCode int
		expBody string
	}

	tt := []test{
		{
			name:    "invalid author id",
			path:    slash + file + slash + api + slash,
			method:  http.MethodPut,
			isOkRes: true,
			req: model.UpdateFileRequest{
				ID:          0,
				Name:        "some",
				Description: "some",
				Size:        1,
				Path:        "some",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				Actual:      true,
				AuthorID:    1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Update", data.req).
					Return(0, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "not correct id",
		},
		{
			name:    "update err",
			path:    slash + file + slash + api + slash,
			method:  http.MethodPut,
			isOkRes: true,
			req: model.UpdateFileRequest{
				ID:          1,
				Name:        "some",
				Description: "some",
				Size:        1,
				Path:        "some",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				Actual:      true,
				AuthorID:    1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Update", data.req).
					Return(0, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash + api + slash,
			method: http.MethodPut,
			req: model.UpdateFileRequest{
				ID:          1,
				Name:        "some",
				Description: "some",
				Size:        1,
				Path:        "some",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				Actual:      true,
				AuthorID:    1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Update", data.req).
					Return(0, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash + api + slash,
			method:  http.MethodPut,
			isOkRes: true,
			req: model.UpdateFileRequest{
				ID:          1,
				Name:        "some",
				Description: "some",
				Size:        1,
				Path:        "some",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				Actual:      true,
				AuthorID:    1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Update", data.req).
					Return(15, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(15),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			assert.Nil(err)

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.ID), body)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)
			if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
			}
			assert.Equal(tc.expBody, r)
		})
	}
}

func TestFile_Delete(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name    string
		path    string
		method  string
		isOkRes bool
		req     model.DeleteFileRequest
		fn      func(fileService *m.File, data test)
		expCode int
		expBody string
	}

	tt := []test{
		{
			name:    "invalid author id",
			path:    slash + file + slash + api + slash,
			method:  http.MethodDelete,
			isOkRes: true,
			req: model.DeleteFileRequest{
				ID: 0,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Delete", data.req).
					Return(0, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "not correct id",
		},
		{
			name:    "delete err",
			path:    slash + file + slash + api + slash,
			method:  http.MethodDelete,
			isOkRes: true,
			req: model.DeleteFileRequest{
				ID: 1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Delete", data.req).
					Return(0, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash + api + slash,
			method: http.MethodDelete,
			req: model.DeleteFileRequest{
				ID: 1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Delete", data.req).
					Return(0, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash + api + slash,
			method:  http.MethodDelete,
			isOkRes: true,
			req: model.DeleteFileRequest{
				ID: 1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("Delete", data.req).
					Return(15, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(15),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.ID), nil)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)
			if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
			}
			assert.Equal(tc.expBody, r)
		})
	}
}

func TestFile_FindByID(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkRes     bool
		isOkMessage bool
		req         model.IDFileRequest
		fn          func(fileService *m.File, data test)
		expCode     int
		expRes      model.File
		message     string
	}

	tt := []test{
		{
			name:        "invalid author id",
			path:        slash + file + slash + api + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.IDFileRequest{
				ID: 0,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByID", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + file + slash + api + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.IDFileRequest{
				ID: 1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByID", data.req).
					Return(&data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash + api + slash,
			method: http.MethodGet,
			req: model.IDFileRequest{
				ID: 1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByID", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash + api + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.IDFileRequest{
				ID: 1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByID", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: model.File{
				ID:          1,
				Name:        "some",
				Description: "some",
				Size:        1,
				Path:        "some",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				Actual:      true,
				AuthorID:    1,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var f model.File
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.ID), nil)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)

			switch {
			case tc.isOkMessage:
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
				assert.Equal(tc.message, r)
			case tc.isOkRes:
				err = json.NewDecoder(res.Body).Decode(&f)
				assert.Nil(err)
				assert.Equal(tc.expRes, f)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestFile_FindByAuthorID(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkRes     bool
		isOkMessage bool
		req         model.AuthorIDFileRequest
		fn          func(fileService *m.File, data test)
		expCode     int
		expRes      []model.File
		message     string
	}

	tt := []test{
		{
			name:        "invalid author id",
			path:        slash + file + slash + api + slash + author + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.AuthorIDFileRequest{
				ID: 0,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByAuthorID", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + file + slash + api + slash + author + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.AuthorIDFileRequest{
				ID: 1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByAuthorID", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash + api + slash + author + slash,
			method: http.MethodGet,
			req: model.AuthorIDFileRequest{
				ID: 1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByAuthorID", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash + api + slash + author + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.AuthorIDFileRequest{
				ID: 1,
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByAuthorID", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Actual:      true,
					AuthorID:    1,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var f []model.File
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.ID), nil)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)

			switch {
			case tc.isOkMessage:
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
				assert.Equal(tc.message, r)
			case tc.isOkRes:
				err = json.NewDecoder(res.Body).Decode(&f)
				assert.Nil(err)
				assert.Equal(tc.expRes, f)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestFile_FindByName(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkRes     bool
		isOkMessage bool
		req         model.NameFileRequest
		fn          func(fileService *m.File, data test)
		expCode     int
		expRes      []model.File
		message     string
	}

	tt := []test{
		{
			name:        "find err",
			path:        slash + file + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.NameFileRequest{
				Name: "some",
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByName", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash,
			method: http.MethodGet,
			req: model.NameFileRequest{
				Name: "some",
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByName", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.NameFileRequest{
				Name: "some",
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindByName", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Actual:      true,
					AuthorID:    1,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var f []model.File
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path+tc.req.Name, nil)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)

			switch {
			case tc.isOkMessage:
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
				assert.Equal(tc.message, r)
			case tc.isOkRes:
				err = json.NewDecoder(res.Body).Decode(&f)
				assert.Nil(err)
				assert.Equal(tc.expRes, f)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestFile_FindActual(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkRes     bool
		isOkMessage bool
		fn          func(fileService *m.File, data test)
		expCode     int
		expRes      []model.File
		message     string
	}

	tt := []test{
		{
			name:        "find err",
			path:        slash + file + slash + actual + slash,
			method:      http.MethodGet,
			isOkMessage: true,

			fn: func(fileService *m.File, data test) {
				fileService.On("FindActual").
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash + actual + slash,
			method: http.MethodGet,
			fn: func(fileService *m.File, data test) {
				fileService.On("FindActual").
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash + actual + slash,
			method:  http.MethodGet,
			isOkRes: true,
			fn: func(fileService *m.File, data test) {
				fileService.On("FindActual").
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Actual:      true,
					AuthorID:    1,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var f []model.File
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path, nil)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)

			switch {
			case tc.isOkMessage:
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
				assert.Equal(tc.message, r)
			case tc.isOkRes:
				err = json.NewDecoder(res.Body).Decode(&f)
				assert.Nil(err)
				assert.Equal(tc.expRes, f)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestFile_FindNotActual(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkRes     bool
		isOkMessage bool
		fn          func(fileService *m.File, data test)
		expCode     int
		expRes      []model.File
		message     string
	}

	tt := []test{
		{
			name:        "find err",
			path:        slash + file + slash + expired + slash,
			method:      http.MethodGet,
			isOkMessage: true,

			fn: func(fileService *m.File, data test) {
				fileService.On("FindNotActual").
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash + expired + slash,
			method: http.MethodGet,
			fn: func(fileService *m.File, data test) {
				fileService.On("FindNotActual").
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash + expired + slash,
			method:  http.MethodGet,
			isOkRes: true,
			fn: func(fileService *m.File, data test) {
				fileService.On("FindNotActual").
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Actual:      false,
					AuthorID:    1,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var f []model.File
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path, nil)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)

			switch {
			case tc.isOkMessage:
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
				assert.Equal(tc.message, r)
			case tc.isOkRes:
				err = json.NewDecoder(res.Body).Decode(&f)
				assert.Nil(err)
				assert.Equal(tc.expRes, f)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestFile_FindAll(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkRes     bool
		isOkMessage bool
		fn          func(fileService *m.File, data test)
		expCode     int
		expRes      []model.File
		message     string
	}

	tt := []test{
		{
			name:        "find err",
			path:        slash + file + slash,
			method:      http.MethodGet,
			isOkMessage: true,

			fn: func(fileService *m.File, data test) {
				fileService.On("FindAll").
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash,
			method: http.MethodGet,
			fn: func(fileService *m.File, data test) {
				fileService.On("FindAll").
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash,
			method:  http.MethodGet,
			isOkRes: true,
			fn: func(fileService *m.File, data test) {
				fileService.On("FindAll").
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Actual:      false,
					AuthorID:    1,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var f []model.File
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path, nil)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)

			switch {
			case tc.isOkMessage:
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
				assert.Equal(tc.message, r)
			case tc.isOkRes:
				err = json.NewDecoder(res.Body).Decode(&f)
				assert.Nil(err)
				assert.Equal(tc.expRes, f)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestFile_FindAddedByPeriod(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkRes     bool
		isOkMessage bool
		req         model.AddedPeriodFileRequest
		fn          func(fileService *m.File, data test)
		expCode     int
		expRes      []model.File
		message     string
	}

	tt := []test{
		{
			name:        "invalid date",
			path:        slash + file + slash + added,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.AddedPeriodFileRequest{
				Start: time.Time{},
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindAddedByPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "start is required",
		},
		{
			name:        "find err",
			path:        slash + file + slash + added,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.AddedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindAddedByPeriod", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash + added,
			method: http.MethodPost,
			req: model.AddedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindAddedByPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash + added,
			method:  http.MethodPost,
			isOkRes: true,
			req: model.AddedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindAddedByPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Actual:      true,
					AuthorID:    1,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var f []model.File
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			assert.Nil(err)

			req, err := http.NewRequest(tc.method, tc.path, body)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)

			switch {
			case tc.isOkMessage:
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
				assert.Equal(tc.message, r)
			case tc.isOkRes:
				err = json.NewDecoder(res.Body).Decode(&f)
				assert.Nil(err)
				assert.Equal(tc.expRes, f)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestFile_FindUpdatedByPeriod(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkRes     bool
		isOkMessage bool
		req         model.UpdatedPeriodFileRequest
		fn          func(fileService *m.File, data test)
		expCode     int
		expRes      []model.File
		message     string
	}

	tt := []test{
		{
			name:        "invalid date",
			path:        slash + file + slash + updated,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.UpdatedPeriodFileRequest{
				Start: time.Time{},
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindUpdatedByPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "start is required",
		},
		{
			name:        "find err",
			path:        slash + file + slash + updated,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.UpdatedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindUpdatedByPeriod", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + file + slash + updated,
			method: http.MethodPost,
			req: model.UpdatedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindUpdatedByPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + file + slash + updated,
			method:  http.MethodPost,
			isOkRes: true,
			req: model.UpdatedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(fileService *m.File, data test) {
				fileService.On("FindUpdatedByPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Actual:      true,
					AuthorID:    1,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var f []model.File
			file := new(m.File)
			testAPI.Services.File = file
			router := newFile(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(file, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			assert.Nil(err)

			req, err := http.NewRequest(tc.method, tc.path, body)
			assert.Nil(err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(tc.expCode, res.Code)

			switch {
			case tc.isOkMessage:
				err = json.NewDecoder(res.Body).Decode(&r)
				assert.Nil(err)
				assert.Equal(tc.message, r)
			case tc.isOkRes:
				err = json.NewDecoder(res.Body).Decode(&f)
				assert.Nil(err)
				assert.Equal(tc.expRes, f)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}
