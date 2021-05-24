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
	"github.com/pkg/errors"
	testAssert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const author = "author"

func TestAuthor_Create(t *testing.T) {
	assert := testAssert.New(t)
	testAPI, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testAPI.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name    string
		path    string
		method  string
		req     model.CreateAuthorRequest
		fn      func(authorService *m.Author, data test)
		expCode int
		expBody string
	}

	tt := []test{
		{
			name:   "invalid user id",
			path:   slash + author + slash + api + slash,
			method: http.MethodPost,
			req: model.CreateAuthorRequest{
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      0,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Create", data.req).
					Return(0, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "not correct user id",
		},
		{
			name:   "create err",
			path:   slash + author + slash + api + slash,
			method: http.MethodPost,
			req: model.CreateAuthorRequest{
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Create", data.req).
					Return(0, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "all ok",
			path:   slash + author + slash + api + slash,
			method: http.MethodPost,
			req: model.CreateAuthorRequest{
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Create", data.req).
					Return(15, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(15),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			author := new(m.Author)
			testAPI.Services.Author = author
			router := newAuthor(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(author, tc)
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

func TestAuthor_Update(t *testing.T) {
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
		req     model.UpdateAuthorRequest
		fn      func(authorService *m.Author, data test)
		expCode int
		expBody string
	}

	tt := []test{
		{
			name:    "invalid user id",
			path:    slash + author + slash + api + slash,
			method:  http.MethodPut,
			isOkRes: true,
			req: model.UpdateAuthorRequest{
				ID:          0,
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Update", data.req).
					Return(0, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "not correct id",
		},
		{
			name:    "update err",
			path:    slash + author + slash + api + slash,
			method:  http.MethodPut,
			isOkRes: true,
			req: model.UpdateAuthorRequest{
				ID:          1,
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Update", data.req).
					Return(0, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + author + slash + api + slash,
			method: http.MethodPut,
			req: model.UpdateAuthorRequest{
				ID:          1,
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Update", data.req).
					Return(0, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + author + slash + api + slash,
			method:  http.MethodPut,
			isOkRes: true,
			req: model.UpdateAuthorRequest{
				ID:          15,
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Update", data.req).
					Return(data.req.ID, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(15),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			author := new(m.Author)
			testAPI.Services.Author = author
			router := newAuthor(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(author, tc)
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

func TestAuthor_Delete(t *testing.T) {
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
		req     model.DeleteAuthorRequest
		fn      func(authorService *m.Author, data test)
		expCode int
		expBody string
	}

	tt := []test{
		{
			name:    "invalid user id",
			path:    slash + author + slash + api + slash,
			method:  http.MethodDelete,
			isOkRes: true,
			req: model.DeleteAuthorRequest{
				ID: 0,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Delete", data.req).
					Return(0, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "not correct id",
		},
		{
			name:    "delete err",
			path:    slash + author + slash + api + slash,
			method:  http.MethodDelete,
			isOkRes: true,
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Delete", data.req).
					Return(0, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + author + slash + api + slash,
			method: http.MethodDelete,
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Delete", data.req).
					Return(0, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + author + slash + api + slash,
			method:  http.MethodDelete,
			isOkRes: true,
			req: model.DeleteAuthorRequest{
				ID: 15,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("Delete", data.req).
					Return(data.req.ID, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(15),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			author := new(m.Author)
			testAPI.Services.Author = author
			router := newAuthor(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(author, tc)
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

func TestAuthor_FindByID(t *testing.T) {
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
		req         model.IDAuthorRequest
		fn          func(authorService *m.Author, data test)
		expCode     int
		expRes      model.Author
		message     string
	}

	tt := []test{
		{
			name:        "invalid user id",
			path:        slash + author + slash + api + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.IDAuthorRequest{
				ID: 0,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByID", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + author + slash + api + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.IDAuthorRequest{
				ID: 1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByID", data.req).
					Return(&data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + author + slash + api + slash,
			method: http.MethodGet,
			req: model.IDAuthorRequest{
				ID: 1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByID", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + author + slash + api + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.IDAuthorRequest{
				ID: 15,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByID", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: model.Author{
				ID:          1,
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var a model.Author
			author := new(m.Author)
			testAPI.Services.Author = author
			router := newAuthor(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(author, tc)
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
				err = json.NewDecoder(res.Body).Decode(&a)
				assert.Nil(err)
				assert.Equal(tc.expRes, a)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestAuthor_FindByUserID(t *testing.T) {
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
		req         model.UserIDAuthorRequest
		fn          func(authorService *m.Author, data test)
		expCode     int
		expRes      model.Author
		message     string
	}

	tt := []test{
		{
			name:        "invalid user id",
			path:        slash + author + slash + api + slash + user + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.UserIDAuthorRequest{
				ID: 0,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByUserID", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + author + slash + api + slash + user + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.UserIDAuthorRequest{
				ID: 1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByUserID", data.req).
					Return(&data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + author + slash + api + slash + user + slash,
			method: http.MethodGet,
			req: model.UserIDAuthorRequest{
				ID: 1,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByUserID", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + author + slash + api + slash + user + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.UserIDAuthorRequest{
				ID: 15,
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByUserID", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: model.Author{
				ID:          1,
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      15,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var a model.Author
			author := new(m.Author)
			testAPI.Services.Author = author
			router := newAuthor(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(author, tc)
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
				err = json.NewDecoder(res.Body).Decode(&a)
				assert.Nil(err)
				assert.Equal(tc.expRes, a)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestAuthor_FindByName(t *testing.T) {
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
		req         model.NameAuthorRequest
		fn          func(authorService *m.Author, data test)
		expCode     int
		expRes      []model.Author
		message     string
	}

	tt := []test{
		{
			name:        "find err",
			path:        slash + author + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.NameAuthorRequest{
				Name: "some",
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByName", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + author + slash,
			method: http.MethodGet,
			req: model.NameAuthorRequest{
				Name: "some",
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByName", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + author + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.NameAuthorRequest{
				Name: "some",
			},
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindByName", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Author{
				{
					ID:          1,
					Name:        "some",
					Age:         1,
					Description: "some",
					UserID:      15,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var a []model.Author
			author := new(m.Author)
			testAPI.Services.Author = author
			router := newAuthor(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(author, tc)
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
				err = json.NewDecoder(res.Body).Decode(&a)
				assert.Nil(err)
				assert.Equal(tc.expRes, a)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}

func TestAuthor_FindAll(t *testing.T) {
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
		fn          func(authorService *m.Author, data test)
		expCode     int
		expRes      []model.Author
		message     string
	}

	tt := []test{
		{
			name:        "find err",
			path:        slash + author + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindAll").
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + author + slash,
			method: http.MethodGet,
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindAll").
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + author + slash,
			method:  http.MethodGet,
			isOkRes: true,
			fn: func(authorService *m.Author, data test) {
				authorService.On("FindAll").
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Author{
				{
					ID:          1,
					Name:        "some",
					Age:         1,
					Description: "some",
					UserID:      15,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var a []model.Author
			author := new(m.Author)
			testAPI.Services.Author = author
			router := newAuthor(testAPI.Services, testAPI.TokenManager)
			if tc.fn != nil {
				tc.fn(author, tc)
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
				err = json.NewDecoder(res.Body).Decode(&a)
				assert.Nil(err)
				assert.Equal(tc.expRes, a)
			default:
				assert.Equal(tc.message, r)
			}
		})
	}
}
