package service

import (
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	m "github.com/JesusG2000/hexsatisfaction/internal/service/mock"
	"github.com/pkg/errors"
	testAssert "github.com/stretchr/testify/assert"
)

func TestFileService_Create(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.CreateFileRequest
		fn     func(file *m.File, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Create errors",
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
			fn: func(file *m.File, data test) {
				file.On("Create", model.File{
					Name:        data.req.Name,
					Description: data.req.Description,
					Size:        data.req.Size,
					Path:        data.req.Path,
					AddDate:     data.req.AddDate,
					UpdateDate:  data.req.UpdateDate,
					Actual:      data.req.Actual,
					AuthorID:    data.req.AuthorID,
				}).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't create file"),
		},
		{
			name: "All ok",
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
			fn: func(file *m.File, data test) {
				file.On("Create", model.File{
					Name:        data.req.Name,
					Description: data.req.Description,
					Size:        data.req.Size,
					Path:        data.req.Path,
					AddDate:     data.req.AddDate,
					UpdateDate:  data.req.UpdateDate,
					Actual:      data.req.Actual,
					AuthorID:    data.req.AuthorID,
				}).
					Return(data.expID, nil)
			},
			expID: 15,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			id, err := service.Create(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestFileService_Update(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UpdateFileRequest
		fn     func(file *m.File, data test)
		expID  int
		expErr error
	}
	tt := []test{
		{
			name: "Update errors",
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
			fn: func(file *m.File, data test) {
				file.On("Update", data.req.ID, model.File{
					Name:        data.req.Name,
					Description: data.req.Description,
					Size:        data.req.Size,
					Path:        data.req.Path,
					AddDate:     data.req.AddDate,
					UpdateDate:  data.req.UpdateDate,
					Actual:      data.req.Actual,
					AuthorID:    data.req.AuthorID,
				}).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't update file"),
		},
		{
			name: "All ok",
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
			fn: func(file *m.File, data test) {
				file.On("Update", data.req.ID, model.File{
					Name:        data.req.Name,
					Description: data.req.Description,
					Size:        data.req.Size,
					Path:        data.req.Path,
					AddDate:     data.req.AddDate,
					UpdateDate:  data.req.UpdateDate,
					Actual:      data.req.Actual,
					AuthorID:    data.req.AuthorID,
				}).
					Return(data.expID, nil)
			},
			expID: 1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			id, err := service.Update(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestFileService_Delete(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name        string
		req         model.DeleteFileRequest
		fn          func(file *m.File, purchase *m.Purchase, comment *m.Comment, data test)
		expPurchase []model.Purchase
		expID       int
		expErr      error
	}
	tt := []test{
		{
			name: "Find purchase errors",
			expPurchase: []model.Purchase{
				{
					ID:     1,
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileID: 1,
				},
			},
			req: model.DeleteFileRequest{
				ID: 1,
			},
			fn: func(file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				purchase.On("FindByFileID", data.req.ID).
					Return(data.expPurchase, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't get purchases"),
		},
		{
			name: "Delete comments errors",
			expPurchase: []model.Purchase{
				{
					ID:     1,
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileID: 1,
				},
			},
			req: model.DeleteFileRequest{
				ID: 1,
			},
			fn: func(file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				purchase.On("FindByFileID", data.req.ID).
					Return(data.expPurchase, nil)
				for _, p := range data.expPurchase {
					comment.On("DeleteByPurchaseID", p.ID).
						Return(0, errors.New(""))
				}
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete comment"),
		},
		{
			name: "Delete purchases errors",
			expPurchase: []model.Purchase{
				{
					ID:     1,
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileID: 1,
				},
			},
			req: model.DeleteFileRequest{
				ID: 1,
			},
			fn: func(file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				purchase.On("FindByFileID", data.req.ID).
					Return(data.expPurchase, nil)
				for _, p := range data.expPurchase {
					comment.On("DeleteByPurchaseID", p.ID).
						Return(p.ID, nil)
				}
				purchase.On("DeleteByFileID", data.req.ID).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete purchases"),
		},
		{
			name: "Delete file errors",
			expPurchase: []model.Purchase{
				{
					ID:     1,
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileID: 1,
				},
			},
			req: model.DeleteFileRequest{
				ID: 1,
			},
			fn: func(file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				purchase.On("FindByFileID", data.req.ID).
					Return(data.expPurchase, nil)
				for _, p := range data.expPurchase {
					comment.On("DeleteByPurchaseID", p.ID).
						Return(p.ID, nil)
				}
				purchase.On("DeleteByFileID", data.req.ID).
					Return(data.expID, nil)
				file.On("Delete", data.expID).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete file"),
		},
		{
			name: "All ok",
			expPurchase: []model.Purchase{
				{
					ID:     1,
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileID: 1,
				},
			},
			req: model.DeleteFileRequest{
				ID: 1,
			},
			fn: func(file *m.File, purchase *m.Purchase, comment *m.Comment, data test) {
				purchase.On("FindByFileID", data.req.ID).
					Return(data.expPurchase, nil)
				for _, p := range data.expPurchase {
					comment.On("DeleteByPurchaseID", p.ID).
						Return(p.ID, nil)
				}
				purchase.On("DeleteByFileID", data.req.ID).
					Return(data.expID, nil)
				file.On("Delete", data.expID).
					Return(data.expID, nil)
			},
			expID: 1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, purchase, comment, tc)
			}
			id, err := service.Delete(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestFileService_FindByID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.IDFileRequest
		fn     func(file *m.File, data test)
		exp    *model.File
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.IDFileRequest{
				ID: 1,
			},
			fn: func(file *m.File, data test) {
				file.On("FindByID", data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find file"),
		},
		{
			name: "All ok",
			req: model.IDFileRequest{
				ID: 1,
			},
			fn: func(file *m.File, data test) {
				file.On("FindByID", data.req.ID).
					Return(data.exp, nil)
			},
			exp: &model.File{
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			f, err := service.FindByID(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, f)
		})
	}
}

func TestFileService_FindByName(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.NameFileRequest
		fn     func(file *m.File, data test)
		exp    []model.File
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.NameFileRequest{
				Name: "some",
			},
			fn: func(file *m.File, data test) {
				file.On("FindByName", data.req.Name).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find files"),
		},
		{
			name: "All ok",
			req: model.NameFileRequest{
				Name: "some",
			},
			fn: func(file *m.File, data test) {
				file.On("FindByName", data.req.Name).
					Return(data.exp, nil)
			},
			exp: []model.File{
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			f, err := service.FindByName(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, f)
		})
	}
}

func TestFileService_FindAll(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		fn     func(file *m.File, data test)
		exp    []model.File
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",

			fn: func(file *m.File, data test) {
				file.On("FindAll").
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find files"),
		},
		{
			name: "All ok",

			fn: func(file *m.File, data test) {
				file.On("FindAll").
					Return(data.exp, nil)
			},
			exp: []model.File{
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			f, err := service.FindAll()
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, f)
		})
	}
}

func TestFileService_FindByAuthorID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.AuthorIDFileRequest
		fn     func(file *m.File, data test)
		exp    []model.File
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.AuthorIDFileRequest{
				ID: 1,
			},
			fn: func(file *m.File, data test) {
				file.On("FindByAuthorID", data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find files"),
		},
		{
			name: "All ok",
			req: model.AuthorIDFileRequest{
				ID: 1,
			},
			fn: func(file *m.File, data test) {
				file.On("FindByAuthorID", data.req.ID).
					Return(data.exp, nil)
			},
			exp: []model.File{
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			f, err := service.FindByAuthorID(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, f)
		})
	}
}

func TestFileService_FindNotActual(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		fn     func(file *m.File, data test)
		exp    []model.File
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",

			fn: func(file *m.File, data test) {
				file.On("FindNotActual").
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find files"),
		},
		{
			name: "All ok",

			fn: func(file *m.File, data test) {
				file.On("FindNotActual").
					Return(data.exp, nil)
			},
			exp: []model.File{
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			f, err := service.FindNotActual()
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, f)
		})
	}
}

func TestFileService_FindActual(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		fn     func(file *m.File, data test)
		exp    []model.File
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",

			fn: func(file *m.File, data test) {
				file.On("FindActual").
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find files"),
		},
		{
			name: "All ok",

			fn: func(file *m.File, data test) {
				file.On("FindActual").
					Return(data.exp, nil)
			},
			exp: []model.File{
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			f, err := service.FindActual()
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, f)
		})
	}
}

func TestFileService_FindAddedByPeriod(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.AddedPeriodFileRequest
		fn     func(file *m.File, data test)
		exp    []model.File
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.AddedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(file *m.File, data test) {
				file.On("FindAddedByPeriod", data.req.Start, data.req.End).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find files"),
		},
		{
			name: "All ok",
			req: model.AddedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(file *m.File, data test) {
				file.On("FindAddedByPeriod", data.req.Start, data.req.End).
					Return(data.exp, nil)
			},
			exp: []model.File{
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			f, err := service.FindAddedByPeriod(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, f)
		})
	}
}

func TestFileService_FindUpdatedByPeriod(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UpdatedPeriodFileRequest
		fn     func(file *m.File, data test)
		exp    []model.File
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.UpdatedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(file *m.File, data test) {
				file.On("FindUpdatedByPeriod", data.req.Start, data.req.End).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find files"),
		},
		{
			name: "All ok",
			req: model.UpdatedPeriodFileRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(file *m.File, data test) {
				file.On("FindUpdatedByPeriod", data.req.Start, data.req.End).
					Return(data.exp, nil)
			},
			exp: []model.File{
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
			file := new(m.File)
			purchase := new(m.Purchase)
			comment := new(m.Comment)
			service := NewFileService(file, purchase, comment)
			if tc.fn != nil {
				tc.fn(file, tc)
			}
			f, err := service.FindUpdatedByPeriod(tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, f)
		})
	}
}
