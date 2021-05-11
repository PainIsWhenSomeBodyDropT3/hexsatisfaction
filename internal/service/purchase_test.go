package service

import (
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestPurchaseService_Create(t *testing.T) {
	type test struct {
		name   string
		req    model.CreatePurchaseRequest
		fn     func(purchase *m.Purchase, data test)
		expId  int
		expErr error
	}
	tt := []test{
		{
			name: "Create errors",
			req: model.CreatePurchaseRequest{
				UserId:   23,
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("Create", model.Purchase{
					UserId:   data.req.UserId,
					Date:     data.req.Date,
					FileName: data.req.FileName,
				}).
					Return(data.expId, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't create purchase"),
		},
		{
			name: "All ok",
			req: model.CreatePurchaseRequest{
				UserId:   23,
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("Create", model.Purchase{
					UserId:   data.req.UserId,
					Date:     data.req.Date,
					FileName: data.req.FileName,
				}).
					Return(data.expId, nil)
			},
			expId: 15,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
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

func TestPurchaseService_Delete(t *testing.T) {
	type test struct {
		name   string
		req    model.DeletePurchaseRequest
		fn     func(purchase *m.Purchase, data test)
		expId  int
		expErr error
	}
	tt := []test{
		{
			name: "Delete errors",
			req: model.DeletePurchaseRequest{
				Id: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("Delete", data.req.Id).
					Return(data.expId, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete purchase"),
		},
		{
			name: "All ok",
			req: model.DeletePurchaseRequest{
				Id: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("Delete", data.req.Id).
					Return(data.expId, nil)
			},
			expId: 15,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			id, err := service.Delete(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expId, id)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindById(t *testing.T) {
	type test struct {
		name        string
		req         model.IdPurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase *model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindById errors",
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindById", 0).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchase"),
		},
		{
			name: "All ok",
			req: model.IdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindById", data.req.Id).
					Return(data.expPurchase, nil)
			},
			expPurchase: &model.Purchase{
				ID:       15,
				UserId:   23,
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindById(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindLastByUserId(t *testing.T) {
	type test struct {
		name        string
		req         model.UserIdPurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase *model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindLastByUserId errors",
			req: model.UserIdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindLastByUserId", data.req.Id).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchase"),
		},
		{
			name: "All ok",
			req: model.UserIdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindLastByUserId", data.req.Id).
					Return(data.expPurchase, nil)
			},
			expPurchase: &model.Purchase{
				ID:       15,
				UserId:   23,
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindLastByUserId(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindAllByUserId(t *testing.T) {
	type test struct {
		name        string
		req         model.UserIdPurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindAllByUserId errors",
			req: model.UserIdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindAllByUserId", data.req.Id).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdPurchaseRequest{
				Id: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindAllByUserId", data.req.Id).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindAllByUserId(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindByUserIdAndPeriod(t *testing.T) {
	type test struct {
		name        string
		req         model.UserIdPeriodPurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByUserIdAndPeriod errors",
			req: model.UserIdPeriodPurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIdAndPeriod", data.req.Id, data.req.Start, data.req.End).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdPeriodPurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIdAndPeriod", data.req.Id, data.req.Start, data.req.End).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 15, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 3, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindByUserIdAndPeriod(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindByUserIdAfterDate(t *testing.T) {
	type test struct {
		name        string
		req         model.UserIdAfterDatePurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByUserIdAfterDate errors",
			req: model.UserIdAfterDatePurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIdAfterDate", data.req.Id, data.req.Start).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdAfterDatePurchaseRequest{
				Id:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIdAfterDate", data.req.Id, data.req.Start).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindByUserIdAfterDate(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindByUserIdBeforeDate(t *testing.T) {
	type test struct {
		name        string
		req         model.UserIdBeforeDatePurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByUserIdBeforeDate errors",
			req: model.UserIdBeforeDatePurchaseRequest{
				Id:  15,
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIdBeforeDate", data.req.Id, data.req.End).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdBeforeDatePurchaseRequest{
				Id:  15,
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIdBeforeDate", data.req.Id, data.req.End).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindByUserIdBeforeDate(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindByUserIdAndFileName(t *testing.T) {
	type test struct {
		name        string
		req         model.UserIdFileNamePurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByUserIdAndFileName errors",
			req: model.UserIdFileNamePurchaseRequest{
				Id:       15,
				FileName: "test",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIdAndFileName", data.req.Id, data.req.FileName).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdFileNamePurchaseRequest{
				Id:       15,
				FileName: "test",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIdAndFileName", data.req.Id, data.req.FileName).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindByUserIdAndFileName(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindLast(t *testing.T) {
	type test struct {
		name        string
		fn          func(purchase *m.Purchase, data test)
		expPurchase *model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindLast errors",
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindLast").
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchase"),
		},
		{
			name: "All ok",
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindLast").
					Return(data.expPurchase, nil)
			},
			expPurchase: &model.Purchase{
				ID:       23,
				UserId:   15,
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindLast()
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindAll(t *testing.T) {
	type test struct {
		name        string
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindAll errors",
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindAll").
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindAll").
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindAll()
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindByPeriod(t *testing.T) {
	type test struct {
		name        string
		req         model.PeriodPurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByPeriod errors",
			req: model.PeriodPurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByPeriod", data.req.Start, data.req.End).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.PeriodPurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByPeriod", data.req.Start, data.req.End).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindByPeriod(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindAfterDate(t *testing.T) {
	type test struct {
		name        string
		req         model.AfterDatePurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindAfterDate errors",
			req: model.AfterDatePurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindAfterDate", data.req.Start).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.AfterDatePurchaseRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindAfterDate", data.req.Start).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindAfterDate(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindBeforeDate(t *testing.T) {
	type test struct {
		name        string
		req         model.BeforeDatePurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindBeforeDate errors",
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindBeforeDate", data.req.End).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.BeforeDatePurchaseRequest{
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindBeforeDate", data.req.End).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindBeforeDate(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestPurchaseService_FindByFileName(t *testing.T) {
	type test struct {
		name        string
		req         model.FileNamePurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByFileName errors",
			req: model.FileNamePurchaseRequest{
				FileName: "test",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByFileName", data.req.FileName).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.FileNamePurchaseRequest{
				FileName: "test",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByFileName", data.req.FileName).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserId:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserId:   15,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase, tc)
			}
			p, err := service.FindByFileName(tc.req)
			if err != nil {
				require.EqualError(t, tc.expErr, err.Error())
				require.Equal(t, tc.expPurchase, p)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
