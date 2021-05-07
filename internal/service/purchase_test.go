package service

import (
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPurchaseService_Create(t *testing.T) {
	id := 23
	tt := []struct {
		name   string
		req    model.CreatePurchaseRequest
		fn     func(purchase *m.Purchase)
		expId  int
		expErr error
	}{
		{
			name: "Create errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("Create", mock.Anything).
					Return(0, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't create purchase"),
		},
		{
			name: "All ok",
			req: model.CreatePurchaseRequest{
				UserId:   id,
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("Create", model.Purchase{
					UserId:   id,
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				}).
					Return(id, nil)
			},
			expId: id,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase)
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
	id := 23
	tt := []struct {
		name   string
		req    model.DeletePurchaseRequest
		fn     func(purchase *m.Purchase)
		expId  int
		expErr error
	}{
		{
			name: "Delete errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("Delete", mock.Anything).
					Return(0, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete purchase"),
		},
		{
			name: "All ok",
			req: model.DeletePurchaseRequest{
				Id: id,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("Delete", id).
					Return(id, nil)
			},
			expId: id,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			purchase := new(m.Purchase)
			service := NewPurchaseService(purchase)
			if tc.fn != nil {
				tc.fn(purchase)
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
	id := 23
	tt := []struct {
		name        string
		req         model.IdPurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase *model.Purchase
		expErr      error
	}{
		{
			name: "FindById errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindById", mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchase"),
		},
		{
			name: "All ok",
			req: model.IdPurchaseRequest{
				Id: id,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindById", id).
					Return(&model.Purchase{
						ID:       id,
						UserId:   id,
						Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
						FileName: "some name",
					}, nil)
			},
			expPurchase: &model.Purchase{
				ID:       id,
				UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	tt := []struct {
		name        string
		req         model.UserIdPurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase *model.Purchase
		expErr      error
	}{
		{
			name: "FindLastByUserId errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindLastByUserId", mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchase"),
		},
		{
			name: "All ok",
			req: model.UserIdPurchaseRequest{
				Id: id,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindLastByUserId", id).
					Return(&model.Purchase{
						ID:       id,
						UserId:   id,
						Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
						FileName: "some name",
					}, nil)
			},
			expPurchase: &model.Purchase{
				ID:       id,
				UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	tt := []struct {
		name        string
		req         model.UserIdPurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindAllByUserId errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindAllByUserId", mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdPurchaseRequest{
				Id: id,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindAllByUserId", id).
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name1",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	start := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	end := time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local)
	tt := []struct {
		name        string
		req         model.UserIdPeriodPurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindByUserIdAndPeriod errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByUserIdAndPeriod", mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdPeriodPurchaseRequest{
				Id:    id,
				Start: start,
				End:   end,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByUserIdAndPeriod", id, start, end).
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name1",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	start := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	tt := []struct {
		name        string
		req         model.UserIdAfterDatePurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindByUserIdAfterDate errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByUserIdAfterDate", mock.Anything, mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdAfterDatePurchaseRequest{
				Id:    id,
				Start: start,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByUserIdAfterDate", id, start).
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name1",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	end := time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local)
	tt := []struct {
		name        string
		req         model.UserIdBeforeDatePurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindByUserIdBeforeDate errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByUserIdBeforeDate", mock.Anything, mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdBeforeDatePurchaseRequest{
				Id:  id,
				End: end,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByUserIdBeforeDate", id, end).
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name1",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	name := "some name"
	tt := []struct {
		name        string
		req         model.UserIdFileNamePurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindByUserIdAndFileName errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByUserIdAndFileName", mock.Anything, mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.UserIdFileNamePurchaseRequest{
				Id:       id,
				FileName: name,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByUserIdAndFileName", id, name).
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	tt := []struct {
		name        string
		fn          func(purchase *m.Purchase)
		expPurchase *model.Purchase
		expErr      error
	}{
		{
			name: "FindLast errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindLast").
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchase"),
		},
		{
			name: "All ok",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindLast").
					Return(&model.Purchase{
						ID:       id,
						UserId:   id,
						Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
						FileName: "some name",
					}, nil)
			},
			expPurchase: &model.Purchase{
				ID:       id,
				UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	tt := []struct {
		name        string
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindAll errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindAll").
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindAll").
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name1",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	start := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	end := time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local)
	tt := []struct {
		name        string
		req         model.PeriodPurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindByPeriod errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByPeriod", mock.Anything, mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.PeriodPurchaseRequest{
				Start: start,
				End:   end,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByPeriod", start, end).
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name1",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	start := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	tt := []struct {
		name        string
		req         model.AfterDatePurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindAfterDate errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindAfterDate", mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.AfterDatePurchaseRequest{
				Start: start,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindAfterDate", start).
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name1",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	end := time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local)
	tt := []struct {
		name        string
		req         model.BeforeDatePurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindBeforeDate errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindBeforeDate", mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.BeforeDatePurchaseRequest{
				End: end,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindBeforeDate", end).
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name1",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
	id := 23
	name := "some name"
	tt := []struct {
		name        string
		req         model.FileNamePurchaseRequest
		fn          func(purchase *m.Purchase)
		expPurchase []model.Purchase
		expErr      error
	}{
		{
			name: "FindByFileName errors",
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByFileName", mock.Anything).
					Return(nil, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find purchases"),
		},
		{
			name: "All ok",
			req: model.FileNamePurchaseRequest{
				FileName: name,
			},
			fn: func(purchase *m.Purchase) {
				purchase.On("FindByFileName", name).
					Return([]model.Purchase{
						{
							ID:       id,
							UserId:   id,
							Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
						{
							ID:       id + 1,
							UserId:   id,
							Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
							FileName: "some name",
						},
					}, nil)
			},
			expPurchase: []model.Purchase{
				{
					ID:       id,
					UserId:   id,
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name",
				},
				{
					ID:       id + 1,
					UserId:   id,
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
				tc.fn(purchase)
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
