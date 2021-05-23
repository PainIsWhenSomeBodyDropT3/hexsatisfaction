package service

import (
	"testing"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	testAssert "github.com/stretchr/testify/assert"
)

func TestAuthorService_Create(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.CreateAuthorRequest
		fn     func(author *m.Author, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Create errors",
			req: model.CreateAuthorRequest{
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(author *m.Author, data test) {
				author.On("Create", model.Author{
					Name:        data.req.Name,
					Age:         data.req.Age,
					Description: data.req.Description,
					UserID:      data.req.UserID,
				}).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't create author"),
		},
		{
			name: "All ok",
			req: model.CreateAuthorRequest{
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(author *m.Author, data test) {
				author.On("Create", model.Author{
					Name:        data.req.Name,
					Age:         data.req.Age,
					Description: data.req.Description,
					UserID:      data.req.UserID,
				}).
					Return(data.expID, nil)
			},
			expID: 15,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			author := new(m.Author)
			service := NewAuthorService(author)
			if tc.fn != nil {
				tc.fn(author, tc)
			}
			id, err := service.Create(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestAuthorService_Update(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UpdateAuthorRequest
		fn     func(author *m.Author, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Update errors",
			req: model.UpdateAuthorRequest{
				ID:          1,
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(author *m.Author, data test) {
				author.On("Update", data.req.ID, model.Author{
					Name:        data.req.Name,
					Age:         data.req.Age,
					Description: data.req.Description,
					UserID:      data.req.UserID,
				}).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't update author"),
		},
		{
			name: "All ok",
			req: model.UpdateAuthorRequest{
				ID:          1,
				Name:        "some",
				Age:         1,
				Description: "some",
				UserID:      1,
			},
			fn: func(author *m.Author, data test) {
				author.On("Update", data.req.ID, model.Author{
					Name:        data.req.Name,
					Age:         data.req.Age,
					Description: data.req.Description,
					UserID:      data.req.UserID,
				}).
					Return(data.expID, nil)
			},
			expID: 1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			author := new(m.Author)
			service := NewAuthorService(author)
			if tc.fn != nil {
				tc.fn(author, tc)
			}
			id, err := service.Update(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestAuthorService_Delete(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.DeleteAuthorRequest
		fn     func(author *m.Author, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Delete errors",
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			fn: func(author *m.Author, data test) {
				author.On("Delete", data.req.ID).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete author"),
		},
		{
			name: "All ok",
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			fn: func(author *m.Author, data test) {
				author.On("Delete", data.req.ID).
					Return(data.expID, nil)
			},
			expID: 1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			author := new(m.Author)
			service := NewAuthorService(author)
			if tc.fn != nil {
				tc.fn(author, tc)
			}
			id, err := service.Delete(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestAuthorService_FindByID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.IDAuthorRequest
		fn     func(author *m.Author, data test)
		exp    *model.Author
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.IDAuthorRequest{
				ID: 1,
			},
			fn: func(author *m.Author, data test) {
				author.On("FindByID", data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find author"),
		},
		{
			name: "All ok",
			req: model.IDAuthorRequest{
				ID: 1,
			},
			fn: func(author *m.Author, data test) {
				author.On("FindByID", data.req.ID).
					Return(data.exp, nil)
			},
			exp: &model.Author{
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
			author := new(m.Author)
			service := NewAuthorService(author)
			if tc.fn != nil {
				tc.fn(author, tc)
			}
			a, err := service.FindByID(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, a)
		})
	}
}

func TestAuthorService_FindByUserID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UserIDAuthorRequest
		fn     func(author *m.Author, data test)
		exp    *model.Author
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.UserIDAuthorRequest{
				ID: 1,
			},
			fn: func(author *m.Author, data test) {
				author.On("FindByUserID", data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find author"),
		},
		{
			name: "All ok",
			req: model.UserIDAuthorRequest{
				ID: 1,
			},
			fn: func(author *m.Author, data test) {
				author.On("FindByUserID", data.req.ID).
					Return(data.exp, nil)
			},
			exp: &model.Author{
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
			author := new(m.Author)
			service := NewAuthorService(author)
			if tc.fn != nil {
				tc.fn(author, tc)
			}
			a, err := service.FindByUserID(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, a)
		})
	}
}

func TestAuthorService_FindByName(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.NameAuthorRequest
		fn     func(author *m.Author, data test)
		exp    []model.Author
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.NameAuthorRequest{
				Name: "some",
			},
			fn: func(author *m.Author, data test) {
				author.On("FindByName", data.req.Name).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find authors"),
		},
		{
			name: "All ok",
			req: model.NameAuthorRequest{
				Name: "some",
			},
			fn: func(author *m.Author, data test) {
				author.On("FindByName", data.req.Name).
					Return(data.exp, nil)
			},
			exp: []model.Author{
				{
					ID:          1,
					Name:        "some",
					Age:         1,
					Description: "some",
					UserID:      1,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			author := new(m.Author)
			service := NewAuthorService(author)
			if tc.fn != nil {
				tc.fn(author, tc)
			}
			a, err := service.FindByName(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, a)
		})
	}
}

func TestAuthorService_FindAll(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		fn     func(author *m.Author, data test)
		exp    []model.Author
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			fn: func(author *m.Author, data test) {
				author.On("FindAll").
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find authors"),
		},
		{
			name: "All ok",
			fn: func(author *m.Author, data test) {
				author.On("FindAll").
					Return(data.exp, nil)
			},
			exp: []model.Author{
				{
					ID:          1,
					Name:        "some",
					Age:         1,
					Description: "some",
					UserID:      1,
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			author := new(m.Author)
			service := NewAuthorService(author)
			if tc.fn != nil {
				tc.fn(author, tc)
			}
			a, err := service.FindAll()
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, a)
		})
	}
}
