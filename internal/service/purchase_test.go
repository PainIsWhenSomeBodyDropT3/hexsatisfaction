package service

import (
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPurchaseService_Create(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name   string
		req    model.CreatePurchaseRequest
		fn     func(purchase *m.Purchase, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Create errors",
			req: model.CreatePurchaseRequest{
				UserID:   23,
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("Create", model.Purchase{
					UserID:   data.req.UserID,
					Date:     data.req.Date,
					FileName: data.req.FileName,
				}).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't create purchase"),
		},
		{
			name: "All ok",
			req: model.CreatePurchaseRequest{
				UserID:   23,
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("Create", model.Purchase{
					UserID:   data.req.UserID,
					Date:     data.req.Date,
					FileName: data.req.FileName,
				}).
					Return(data.expID, nil)
			},
			expID: 15,
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
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expID, id)
		})
	}
}

func TestPurchaseService_Delete(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name   string
		req    model.DeletePurchaseRequest
		fn     func(purchase *m.Purchase, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Delete errors",
			req: model.DeletePurchaseRequest{
				ID: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("Delete", data.req.ID).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete purchase"),
		},
		{
			name: "All ok",
			req: model.DeletePurchaseRequest{
				ID: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("Delete", data.req.ID).
					Return(data.expID, nil)
			},
			expID: 15,
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
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expID, id)
		})
	}
}

func TestPurchaseService_FindById(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name        string
		req         model.IDPurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase *model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByID errors",
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByID", 0).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchase"),
		},
		{
			name: "All ok",
			req: model.IDPurchaseRequest{
				ID: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByID", data.req.ID).
					Return(data.expPurchase, nil)
			},
			expPurchase: &model.Purchase{
				ID:       15,
				UserID:   23,
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
			p, err := service.FindByID(tc.req)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindLastByUserId(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name        string
		req         model.UserIDPurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase *model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindLastByUserID errors",
			req: model.UserIDPurchaseRequest{
				ID: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindLastByUserID", data.req.ID).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchase"),
		},
		{
			name: "All ok",
			req: model.UserIDPurchaseRequest{
				ID: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindLastByUserID", data.req.ID).
					Return(data.expPurchase, nil)
			},
			expPurchase: &model.Purchase{
				ID:       15,
				UserID:   23,
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
			p, err := service.FindLastByUserID(tc.req)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindAllByUserId(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name        string
		req         model.UserIDPurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindAllByUserID errors",
			req: model.UserIDPurchaseRequest{
				ID: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindAllByUserID", data.req.ID).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIDPurchaseRequest{
				ID: 15,
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindAllByUserID", data.req.ID).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserID:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
			p, err := service.FindAllByUserID(tc.req)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindByUserIdAndPeriod(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name        string
		req         model.UserIDPeriodPurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByUserIDAndPeriod errors",
			req: model.UserIDPeriodPurchaseRequest{
				ID:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIDAndPeriod", data.req.ID, data.req.Start, data.req.End).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIDPeriodPurchaseRequest{
				ID:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIDAndPeriod", data.req.ID, data.req.Start, data.req.End).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserID:   15,
					Date:     time.Date(2009, time.November, 15, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
			p, err := service.FindByUserIDAndPeriod(tc.req)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindByUserIdAfterDate(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name        string
		req         model.UserIDAfterDatePurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByUserIDAfterDate errors",
			req: model.UserIDAfterDatePurchaseRequest{
				ID:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIDAfterDate", data.req.ID, data.req.Start).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIDAfterDatePurchaseRequest{
				ID:    15,
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIDAfterDate", data.req.ID, data.req.Start).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserID:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
			p, err := service.FindByUserIDAfterDate(tc.req)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindByUserIdBeforeDate(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name        string
		req         model.UserIDBeforeDatePurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByUserIDBeforeDate errors",
			req: model.UserIDBeforeDatePurchaseRequest{
				ID:  15,
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIDBeforeDate", data.req.ID, data.req.End).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIDBeforeDatePurchaseRequest{
				ID:  15,
				End: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIDBeforeDate", data.req.ID, data.req.End).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserID:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
			p, err := service.FindByUserIDBeforeDate(tc.req)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindByUserIdAndFileName(t *testing.T) {
	a := assert.New(t)
	type test struct {
		name        string
		req         model.UserIDFileNamePurchaseRequest
		fn          func(purchase *m.Purchase, data test)
		expPurchase []model.Purchase
		expErr      error
	}
	tt := []test{
		{
			name: "FindByUserIDAndFileName errors",
			req: model.UserIDFileNamePurchaseRequest{
				ID:       15,
				FileName: "test",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIDAndFileName", data.req.ID, data.req.FileName).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIDFileNamePurchaseRequest{
				ID:       15,
				FileName: "test",
			},
			fn: func(purchase *m.Purchase, data test) {
				purchase.On("FindByUserIDAndFileName", data.req.ID, data.req.FileName).
					Return(data.expPurchase, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       23,
					UserID:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
			p, err := service.FindByUserIDAndFileName(tc.req)
			if err != nil {
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindLast(t *testing.T) {
	a := assert.New(t)
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
				UserID:   15,
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
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindAll(t *testing.T) {
	a := assert.New(t)
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
					UserID:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindByPeriod(t *testing.T) {
	a := assert.New(t)
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
					UserID:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindAfterDate(t *testing.T) {
	a := assert.New(t)
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
					UserID:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindBeforeDate(t *testing.T) {
	a := assert.New(t)
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
					UserID:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}

func TestPurchaseService_FindByFileName(t *testing.T) {
	a := assert.New(t)
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
					UserID:   15,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       24,
					UserID:   15,
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
				a.Equal(tc.expErr.Error(), err.Error())
			}
			a.Equal(tc.expPurchase, p)
		})
	}
}
