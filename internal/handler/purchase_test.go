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
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	purchase = "purchase"
	last     = "last"
	period   = "period"
	after    = "after"
	before   = "before"
	file     = "file"
)

func TestPurchase_Create(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name    string
		path    string
		method  string
		req     model.CreatePurchaseRequest
		fn      func(purchaseService *m.Purchase, data test)
		expCode int
		expBody string
	}

	tt := []test{
		{
			name:   "invalid user id",
			path:   slash + purchase + slash + api + slash,
			method: http.MethodPost,
			req: model.CreatePurchaseRequest{
				UserId:   0,
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("Create", data.req).
					Return(0, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "not correct user id",
		},
		{
			name:   "create err",
			path:   slash + purchase + slash + api + slash,
			method: http.MethodPost,
			req: model.CreatePurchaseRequest{
				UserId:   23,
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("Create", data.req).
					Return(0, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "all ok",
			path:   slash + purchase + slash + api + slash,
			method: http.MethodPost,
			req: model.CreatePurchaseRequest{
				UserId:   23,
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("Create", data.req).
					Return(15, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(15),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path, body)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			err = json.NewDecoder(res.Body).Decode(&r)
			require.NoError(t, err)
			require.Equal(t, tc.expBody, r)
		})
	}
}

func TestPurchase_Delete(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name    string
		path    string
		method  string
		isOkRes bool
		req     model.DeletePurchaseRequest
		fn      func(purchaseService *m.Purchase, data test)
		expCode int
		expBody string
	}

	tt := []test{
		{
			name:    "invalid  id",
			path:    slash + purchase + slash + api + slash,
			method:  http.MethodDelete,
			isOkRes: true,
			req: model.DeletePurchaseRequest{
				Id: 0,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("Delete", data.req).
					Return(0, nil)
			},
			expCode: http.StatusBadRequest,
			expBody: "not correct id",
		},
		{
			name:    "delete err",
			path:    slash + purchase + slash + api + slash,
			method:  http.MethodDelete,
			isOkRes: true,
			req: model.DeletePurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("Delete", data.req).
					Return(0, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash,
			method: http.MethodDelete,
			req: model.DeletePurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("Delete", data.req).
					Return(0, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash,
			method:  http.MethodDelete,
			isOkRes: true,
			req: model.DeletePurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("Delete", data.req).
					Return(15, nil)
			},
			expCode: http.StatusOK,
			expBody: strconv.Itoa(15),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.Id), nil)
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

func TestPurchase_FindById(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.IdPurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid  id",
			path:        slash + purchase + slash + api + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.IdPurchaseRequest{
				Id: 0,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindById", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.IdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindById", data.req).
					Return(&data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash,
			method: http.MethodGet,
			req: model.IdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindById", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.IdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindById", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: model.Purchase{
				ID:       15,
				UserId:   15,
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "test",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.Id), nil)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindLastByUserId(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.UserIdPurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid  id",
			path:        slash + purchase + slash + api + slash + last + slash + user + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.UserIdPurchaseRequest{
				Id: 0,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindLastByUserId", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + last + slash + user + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.UserIdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindLastByUserId", data.req).
					Return(&data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + last + slash + user + slash,
			method: http.MethodGet,
			req: model.UserIdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindLastByUserId", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + last + slash + user + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.IdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindLastByUserId", data.req).
					Return(&data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: model.Purchase{
				ID:       15,
				UserId:   15,
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "test",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.Id), nil)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindAllByUserId(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.UserIdPurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid  id",
			path:        slash + purchase + slash + api + slash + user + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.UserIdPurchaseRequest{
				Id: 0,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAllByUserId", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + user + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.UserIdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAllByUserId", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + user + slash,
			method: http.MethodGet,
			req: model.UserIdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAllByUserId", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + user + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.IdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAllByUserId", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.Id), nil)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindByUserIdAndPeriod(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.UserIdPeriodPurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid  id",
			path:        slash + purchase + slash + api + slash + period + slash + user + slash,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.UserIdPeriodPurchaseRequest{
				Id:    0,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAndPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + period + slash + user + slash,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.UserIdPeriodPurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAndPeriod", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + period + slash + user + slash,
			method: http.MethodPost,
			req: model.UserIdPeriodPurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAndPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + period + slash + user + slash,
			method:  http.MethodPost,
			isOkRes: true,
			req: model.UserIdPeriodPurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAndPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.Id), body)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindByUserIdAfterDate(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.UserIdAfterDatePurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid  id",
			path:        slash + purchase + slash + api + slash + after + slash + user + slash,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.UserIdAfterDatePurchaseRequest{
				Id:    0,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAfterDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + after + slash + user + slash,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.UserIdAfterDatePurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAfterDate", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + after + slash + user + slash,
			method: http.MethodPost,
			req: model.UserIdAfterDatePurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAfterDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + after + slash + user + slash,
			method:  http.MethodPost,
			isOkRes: true,
			req: model.UserIdAfterDatePurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAfterDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.Id), body)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindByUserIdBeforeDate(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.UserIdBeforeDatePurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid  id",
			path:        slash + purchase + slash + api + slash + before + slash + user + slash,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.UserIdBeforeDatePurchaseRequest{
				Id:  0,
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdBeforeDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + before + slash + user + slash,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.UserIdBeforeDatePurchaseRequest{
				Id:  15,
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdBeforeDate", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + before + slash + user + slash,
			method: http.MethodPost,
			req: model.UserIdBeforeDatePurchaseRequest{
				Id:  15,
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdBeforeDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + before + slash + user + slash,
			method:  http.MethodPost,
			isOkRes: true,
			req: model.UserIdBeforeDatePurchaseRequest{
				Id:  15,
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdBeforeDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path+strconv.Itoa(tc.req.Id), body)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindByUserIdAndFileName(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.UserIdFileNamePurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid  id",
			path:        slash + purchase + slash + api + slash + user + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.UserIdFileNamePurchaseRequest{
				Id:       0,
				FileName: "test",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAndFileName", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "not correct id",
		},
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + user + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.UserIdFileNamePurchaseRequest{
				Id:       15,
				FileName: "test",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAndFileName", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + user + slash,
			method: http.MethodGet,
			req: model.UserIdFileNamePurchaseRequest{
				Id:       15,
				FileName: "test",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAndFileName", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + user + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.UserIdFileNamePurchaseRequest{
				Id:       15,
				FileName: "test",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByUserIdAndFileName", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			fullPath := tc.path + strconv.Itoa(tc.req.Id) + slash + file + slash + tc.req.FileName
			req, err := http.NewRequest(tc.method, fullPath, nil)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindLast(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + last + slash,
			method:      http.MethodGet,
			isOkMessage: true,

			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindLast").
					Return(&data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + last + slash,
			method: http.MethodGet,

			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindLast").
					Return(&data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + last + slash,
			method:  http.MethodGet,
			isOkRes: true,

			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindLast").
					Return(&data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: model.Purchase{
				ID:       15,
				UserId:   15,
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "test",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path, nil)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindAll(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAll").
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash,
			method: http.MethodGet,
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAll").
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash,
			method:  http.MethodGet,
			isOkRes: true,
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAll").
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path, nil)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindByPeriod(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.PeriodPurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid  start date",
			path:        slash + purchase + slash + api + slash + period,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.PeriodPurchaseRequest{
				Start: time.Time{},
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "start date is required",
		},
		{
			name:   "find err",
			path:   slash + purchase + slash + api + slash + period,
			method: http.MethodPost,
			req: model.PeriodPurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByPeriod", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + period,
			method: http.MethodPost,
			req: model.PeriodPurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + period,
			method:  http.MethodPost,
			isOkRes: true,
			req: model.PeriodPurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByPeriod", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path, body)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindAfterDate(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.AfterDatePurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid start date",
			path:        slash + purchase + slash + api + slash + after,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.AfterDatePurchaseRequest{
				Start: time.Time{},
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAfterDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "start date is required",
		},
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + after,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.AfterDatePurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAfterDate", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + after,
			method: http.MethodPost,
			req: model.AfterDatePurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAfterDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + after,
			method:  http.MethodPost,
			isOkRes: true,
			req: model.AfterDatePurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindAfterDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path, body)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindBeforeDate(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.BeforeDatePurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{
		{
			name:        "invalid  end date",
			path:        slash + purchase + slash + api + slash + before,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.BeforeDatePurchaseRequest{
				End: time.Time{},
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindBeforeDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusBadRequest,
			message: "end date is required",
		},
		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + before,
			method:      http.MethodPost,
			isOkMessage: true,
			req: model.BeforeDatePurchaseRequest{
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindBeforeDate", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + before,
			method: http.MethodPost,
			req: model.BeforeDatePurchaseRequest{
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindBeforeDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + before,
			method:  http.MethodPost,
			isOkRes: true,
			req: model.BeforeDatePurchaseRequest{
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindBeforeDate", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&tc.req)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.path, body)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}

func TestPurchase_FindByFileName(t *testing.T) {
	testApi, err := service.InitTest4Mock()
	require.NoError(t, err)
	token, err := testApi.TokenManager.NewJWT(mock.Anything)
	require.NoError(t, err)

	type test struct {
		name        string
		path        string
		method      string
		isOkMessage bool
		isOkRes     bool
		req         model.FileNamePurchaseRequest
		fn          func(purchaseService *m.Purchase, data test)
		expCode     int
		expRes      []model.Purchase
		message     string
	}

	tt := []test{

		{
			name:        "find err",
			path:        slash + purchase + slash + api + slash + file + slash,
			method:      http.MethodGet,
			isOkMessage: true,
			req: model.FileNamePurchaseRequest{
				FileName: "test",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByFileName", data.req).
					Return(data.expRes, errors.New(""))
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "not found",
			path:   slash + purchase + slash + api + slash + file + slash,
			method: http.MethodGet,
			req: model.FileNamePurchaseRequest{
				FileName: "test",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByFileName", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusNotFound,
		},
		{
			name:    "all ok",
			path:    slash + purchase + slash + api + slash + file + slash,
			method:  http.MethodGet,
			isOkRes: true,
			req: model.FileNamePurchaseRequest{
				FileName: "test",
			},
			fn: func(purchaseService *m.Purchase, data test) {
				purchaseService.On("FindByFileName", data.req).
					Return(data.expRes, nil)
			},
			expCode: http.StatusOK,
			expRes: []model.Purchase{
				{
					ID:       15,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
				{
					ID:       16,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "test",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var r string
			var p []model.Purchase
			purchaseService := new(m.Purchase)
			testApi.Services.Purchase = purchaseService
			router := newPurchase(testApi.Services, testApi.TokenManager)
			if tc.fn != nil {
				tc.fn(purchaseService, tc)
			}

			req, err := http.NewRequest(tc.method, tc.path+tc.req.FileName, nil)
			require.NoError(t, err)

			req.Header.Set(authorizationHeader, "Bearer "+token)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			require.Equal(t, tc.expCode, res.Code)

			if tc.isOkMessage {
				err = json.NewDecoder(res.Body).Decode(&r)
				require.NoError(t, err)
				require.Equal(t, tc.message, r)
			} else if tc.isOkRes {
				err = json.NewDecoder(res.Body).Decode(&p)
				require.NoError(t, err)
				require.Equal(t, tc.expRes, p)
			} else {
				require.Equal(t, tc.message, r)
			}
		})
	}
}
