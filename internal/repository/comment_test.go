package repository

import (
	"database/sql"
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	testAssert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func deleteCommentData(assertions *testAssert.Assertions, db *sql.DB) {
	_, err := db.Exec("DELETE FROM comment")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM purchase")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM file")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM author")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM users")
	assertions.Nil(err)
}

func TestCommentRepo_Create(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comment  model.Comment
	}{
		{
			name: "all ok",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
			comment: model.Comment{
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Text: "some text",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			tc.comment.UserID = userID
			tc.comment.PurchaseID = purchaseID
			id, err := repos.Comment.Create(tc.comment)
			assert.Nil(err)
			assert.NotZero(id)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_Delete(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		isOk     bool
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comment  model.Comment
	}{
		{
			name: "not found",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
			comment: model.Comment{
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Text: "some text",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var commentID int
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			if tc.isOk {
				tc.comment.UserID = userID
				tc.comment.PurchaseID = purchaseID
				commentID, err = repos.Comment.Create(tc.comment)
				assert.Nil(err)
			}

			id, err := repos.Comment.Delete(commentID)
			assert.Nil(err)
			assert.Equal(commentID, id)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_DeleteByPurchaseID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		isOk     bool
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comment  model.Comment
	}{
		{
			name: "not found",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
			comment: model.Comment{
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Text: "some text",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var purchaseID int
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			if tc.isOk {
				tc.purchase.UserID = userID
				tc.purchase.FileID = fileID
				tc.purchase.FileID = fileID
				purchaseID, err = repos.Purchase.Create(tc.purchase)
				assert.Nil(err)

				tc.comment.UserID = userID
				tc.comment.PurchaseID = purchaseID
				_, err = repos.Comment.Create(tc.comment)
				assert.Nil(err)
				_, err = repos.Comment.Create(tc.comment)
				assert.Nil(err)
			}

			id, err := repos.Comment.DeleteByPurchaseID(purchaseID)
			assert.Nil(err)
			assert.Equal(purchaseID, id)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_Update(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		isOk     bool
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comment  model.Comment
		update   model.Comment
	}{
		{
			name: "not found",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
			update: model.Comment{
				Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
				Text: "changed text",
			},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
			comment: model.Comment{
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Text: "some text",
			},
			update: model.Comment{
				Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
				Text: "changed text",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var commentID int
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			if tc.isOk {
				tc.comment.UserID = userID
				tc.comment.PurchaseID = purchaseID
				commentID, err = repos.Comment.Create(tc.comment)
				assert.Nil(err)
			}

			tc.update.UserID = userID
			tc.update.PurchaseID = purchaseID
			id, err := repos.Comment.Update(commentID, tc.update)
			assert.Nil(err)
			assert.Equal(commentID, id)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindById(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comment  model.Comment
		exp      *model.Comment
	}{
		{
			name: "find err",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
			exp: &model.Comment{},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileID: 1,
			},
			comment: model.Comment{
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Text: "some text",
			},
			exp: &model.Comment{
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Text: "some text",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var id int
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			if tc.isOk {
				tc.comment.UserID = userID
				tc.comment.PurchaseID = purchaseID
				id, err = repos.Comment.Create(tc.comment)
				assert.Nil(err)
				tc.exp.UserID = userID
				tc.exp.PurchaseID = purchaseID
			}
			p, err := repos.Comment.FindByID(id)
			assert.Nil(err)
			tc.exp.Date = p.Date
			tc.exp.ID = p.ID
			assert.Equal(tc.exp, p)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindAllByUserID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comments []model.Comment
		exp      []model.Comment
	}{
		{
			name: "find err",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileID: 1,
			},
			comments: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
			exp: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindAllByUserID(userID)
			assert.Nil(err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(tc.exp, c)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindByPurchaseID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comments []model.Comment
		exp      []model.Comment
	}{
		{
			name: "find err",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileID: 1,
			},
			comments: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
			exp: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindByPurchaseID(purchaseID)
			assert.Nil(err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(tc.exp, c)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindByUserIDAndPurchaseID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comments []model.Comment
		exp      []model.Comment
	}{
		{
			name: "find err",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileID: 1,
			},
			comments: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
			exp: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindByUserIDAndPurchaseID(userID, purchaseID)
			assert.Nil(err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(tc.exp, c)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindAll(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comments []model.Comment
		exp      []model.Comment
	}{
		{
			name: "find err",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileID: 1,
			},
			comments: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
			exp: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindAll()
			assert.Nil(err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(tc.exp, c)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindByText(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comments []model.Comment
		text     string
		exp      []model.Comment
	}{
		{
			name: "find err",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileID: 1,
			},
			comments: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
			text: "text2",
			exp: []model.Comment{
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindByText(tc.text)
			assert.Nil(err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(tc.exp, c)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindByPeriod(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
		comments []model.Comment
		start    time.Time
		end      time.Time
		exp      []model.Comment
	}{
		{
			name: "find err",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileID: 1,
			},
			start: time.Date(2009, time.November, 15, 23, 0, 0, 0, time.Local),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			author: model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
			file: model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
			purchase: model.Purchase{
				Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileID: 1,
			},
			comments: []model.Comment{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text1",
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
			start: time.Date(2009, time.November, 15, 23, 0, 0, 0, time.Local),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			exp: []model.Comment{
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
					Text: "some text2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteCommentData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			tc.purchase.UserID = userID
			tc.purchase.FileID = fileID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindByPeriod(tc.start, tc.end)
			assert.Nil(err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(tc.exp, c)

			deleteCommentData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}
