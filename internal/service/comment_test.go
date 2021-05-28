package service

import (
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	testAssert "github.com/stretchr/testify/assert"
)

func TestCommentService_Create(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.CreateCommentRequest
		fn     func(comment *m.Comment, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Create errors",
			req: model.CreateCommentRequest{
				UserID:     23,
				PurchaseID: 23,
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("Create", model.Comment{
					UserID:     data.req.UserID,
					PurchaseID: data.req.PurchaseID,
					Date:       data.req.Date,
					Text:       data.req.Text,
				}).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't create comment"),
		},
		{
			name: "All ok",
			req: model.CreateCommentRequest{
				UserID:     23,
				PurchaseID: 23,
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("Create", model.Comment{
					UserID:     data.req.UserID,
					PurchaseID: data.req.PurchaseID,
					Date:       data.req.Date,
					Text:       data.req.Text,
				}).
					Return(data.expID, nil)
			},
			expID: 15,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			id, err := service.Create(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestCommentService_Update(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UpdateCommentRequest
		fn     func(comment *m.Comment, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Update errors",
			req: model.UpdateCommentRequest{
				ID:         15,
				UserID:     23,
				PurchaseID: 23,
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("Update", data.req.ID, model.Comment{
					UserID:     data.req.UserID,
					PurchaseID: data.req.PurchaseID,
					Date:       data.req.Date,
					Text:       data.req.Text,
				}).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't update comment"),
		},
		{
			name: "All ok",
			req: model.UpdateCommentRequest{
				ID:         15,
				UserID:     23,
				PurchaseID: 23,
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("Update", data.req.ID, model.Comment{
					UserID:     data.req.UserID,
					PurchaseID: data.req.PurchaseID,
					Date:       data.req.Date,
					Text:       data.req.Text,
				}).
					Return(data.expID, nil)
			},
			expID: 15,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			id, err := service.Update(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestCommentService_Delete(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.DeleteCommentRequest
		fn     func(comment *m.Comment, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Delete errors",
			req: model.DeleteCommentRequest{
				ID: 15,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("Delete", data.req.ID).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete comment"),
		},
		{
			name: "All ok",
			req: model.DeleteCommentRequest{
				ID: 15,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("Delete", data.req.ID).
					Return(data.expID, nil)
			},
			expID: 15,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			id, err := service.Delete(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestCommentService_FindByID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.IDCommentRequest
		fn     func(comment *m.Comment, data test)
		exp    *model.Comment
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.IDCommentRequest{
				ID: 15,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByID", data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comment"),
		},
		{
			name: "All ok",
			req: model.IDCommentRequest{
				ID: 15,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByID", data.req.ID).
					Return(data.exp, nil)
			},
			exp: &model.Comment{
				ID:         15,
				UserID:     23,
				PurchaseID: 23,
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindByID(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindAllByUserID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UserIDCommentRequest
		fn     func(comment *m.Comment, data test)
		exp    []model.Comment
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.UserIDCommentRequest{
				ID: 23,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindAllByUserID", data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.UserIDCommentRequest{
				ID: 23,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindAllByUserID", data.req.ID).
					Return(data.exp, nil)
			},
			exp: []model.Comment{
				{
					ID:         15,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         16,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindAllByUserID(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindAllByPurchaseID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.PurchaseIDCommentRequest
		fn     func(comment *m.Comment, data test)
		exp    []model.Comment
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.PurchaseIDCommentRequest{
				ID: 23,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByPurchaseID", data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.PurchaseIDCommentRequest{
				ID: 23,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByPurchaseID", data.req.ID).
					Return(data.exp, nil)
			},
			exp: []model.Comment{
				{
					ID:         15,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         16,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindByPurchaseID(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindByUserIDAndPurchaseID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UserPurchaseIDCommentRequest
		fn     func(comment *m.Comment, data test)
		exp    []model.Comment
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.UserPurchaseIDCommentRequest{
				UserID:     23,
				PurchaseID: 23,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByUserIDAndPurchaseID", data.req.UserID, data.req.PurchaseID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.UserPurchaseIDCommentRequest{
				UserID:     23,
				PurchaseID: 23,
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByUserIDAndPurchaseID", data.req.UserID, data.req.PurchaseID).
					Return(data.exp, nil)
			},
			exp: []model.Comment{
				{
					ID:         15,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         16,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindByUserIDAndPurchaseID(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindAll(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		fn     func(comment *m.Comment, data test)
		exp    []model.Comment
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",

			fn: func(comment *m.Comment, data test) {
				comment.On("FindAll").
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			fn: func(comment *m.Comment, data test) {
				comment.On("FindAll").
					Return(data.exp, nil)
			},
			exp: []model.Comment{
				{
					ID:         15,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         16,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindAll()
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindByText(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.TextCommentRequest
		fn     func(comment *m.Comment, data test)
		exp    []model.Comment
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.TextCommentRequest{
				Text: "some",
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByText", data.req.Text).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.TextCommentRequest{
				Text: "some",
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByText", data.req.Text).
					Return(data.exp, nil)
			},
			exp: []model.Comment{
				{
					ID:         15,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         16,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindByText(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindByPeriod(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.PeriodCommentRequest
		fn     func(comment *m.Comment, data test)
		exp    []model.Comment
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.PeriodCommentRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByPeriod", data.req.Start, data.req.End).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.PeriodCommentRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByPeriod", data.req.Start, data.req.End).
					Return(data.exp, nil)
			},
			exp: []model.Comment{
				{
					ID:         15,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         16,
					UserID:     23,
					PurchaseID: 23,
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			service := NewCommentService(comment)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindByPeriod(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}
