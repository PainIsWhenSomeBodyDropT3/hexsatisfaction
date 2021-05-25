package service

import (
	"testing"
	"time"

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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewAuthorService(author, file, purchase, comment)
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewAuthorService(author, file, purchase, comment)
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
		name        string
		req         model.DeleteAuthorRequest
		fn          func(author *m.Author, file *m.File, purchase *m.Purchase, comment *m.Comment, data test)
		expFile     []model.File
		expPurchase []model.Purchase
		expID       int
		expErr      error
	}
	tt := []test{
		{
			name: "Find files errors",
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			fn: func(author *m.Author, file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				file.On("FindByAuthorID", data.req.ID).
					Return(data.expFile, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't get files"),
		},
		{
			name: "Find purchases errors",
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			expFile: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      true,
					AuthorID:    1,
				},
			},
			fn: func(author *m.Author, file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				file.On("FindByAuthorID", data.req.ID).
					Return(data.expFile, nil)
				for _, f := range data.expFile {
					purchase.On("FindByFileID", f.ID).
						Return(data.expPurchase, errors.New(""))
				}
			},
			expErr: errors.Wrap(errors.New(""), "couldn't get purchases"),
		},
		{
			name: "Delete comments errors",
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			expFile: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      true,
					AuthorID:    1,
				},
			},
			expPurchase: []model.Purchase{
				{
					ID:     1,
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					FileID: 1,
				},
			},
			fn: func(author *m.Author, file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				file.On("FindByAuthorID", data.req.ID).
					Return(data.expFile, nil)
				for _, f := range data.expFile {
					purchase.On("FindByFileID", f.ID).
						Return(data.expPurchase, nil)
					for _, p := range data.expPurchase {
						comment.On("DeleteByPurchaseID", p.ID).
							Return(0, errors.New(""))
					}
				}
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete comments"),
		},
		{
			name: "Delete files errors",
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			expFile: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      true,
					AuthorID:    1,
				},
			},
			expPurchase: []model.Purchase{
				{
					ID:     1,
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					FileID: 1,
				},
			},
			fn: func(author *m.Author, file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				file.On("FindByAuthorID", data.req.ID).
					Return(data.expFile, nil)
				for _, f := range data.expFile {
					purchase.On("FindByFileID", f.ID).
						Return(data.expPurchase, nil)
					for _, p := range data.expPurchase {
						comment.On("DeleteByPurchaseID", p.ID).
							Return(0, nil)
					}
				}
				file.On("DeleteByAuthorID", data.req.ID).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete files"),
		},
		{
			name: "Delete author errors",
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			expFile: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      true,
					AuthorID:    1,
				},
			},
			expPurchase: []model.Purchase{
				{
					ID:     1,
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					FileID: 1,
				},
			},
			fn: func(author *m.Author, file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				file.On("FindByAuthorID", data.req.ID).
					Return(data.expFile, nil)
				for _, f := range data.expFile {
					purchase.On("FindByFileID", f.ID).
						Return(data.expPurchase, nil)
					for _, p := range data.expPurchase {
						comment.On("DeleteByPurchaseID", p.ID).
							Return(0, nil)
					}
				}
				file.On("DeleteByAuthorID", data.req.ID).
					Return(data.expID, nil)
				author.On("Delete", data.expID).
					Return(data.expID, nil)
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete author"),
		},
		{
			name: "All ok",
			req: model.DeleteAuthorRequest{
				ID: 1,
			},
			expFile: []model.File{
				{
					ID:          1,
					Name:        "some",
					Description: "some",
					Size:        1,
					Path:        "some",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      true,
					AuthorID:    1,
				},
			},
			expPurchase: []model.Purchase{
				{
					ID:     1,
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					FileID: 1,
				},
			},
			fn: func(author *m.Author, file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				file.On("FindByAuthorID", data.req.ID).
					Return(data.expFile, nil)
				for _, f := range data.expFile {
					purchase.On("FindByFileID", f.ID).
						Return(data.expPurchase, nil)
					for _, p := range data.expPurchase {
						comment.On("DeleteByPurchaseID", p.ID).
							Return(0, nil)
					}
				}
				file.On("DeleteByAuthorID", data.req.ID).
					Return(data.expID, nil)
				author.On("Delete", data.expID).
					Return(data.expID, nil)
			},
			expID: 1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			author := new(m.Author)
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewAuthorService(author, file, purchase, comment)
			if tc.fn != nil {
				tc.fn(author, file, purchase, comment, tc)
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewAuthorService(author, file, purchase, comment)
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewAuthorService(author, file, purchase, comment)
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewAuthorService(author, file, purchase, comment)
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewAuthorService(author, file, purchase, comment)
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
