package repository

import (
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommentRepo_Create(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		user     model.User
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
			},
			comment: model.Comment{
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Text: "some text",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			tc.comment.UserID = userID
			tc.comment.PurchaseID = purchaseID
			id, err := repos.Comment.Create(tc.comment)
			assert.Nil(t, err)
			assert.NotZero(t, id)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_Delete(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		user     model.User
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
			},
			comment: model.Comment{
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Text: "some text",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			tc.comment.UserID = userID
			tc.comment.PurchaseID = purchaseID
			commentID, err := repos.Comment.Create(tc.comment)
			assert.Nil(t, err)

			id, err := repos.Comment.Delete(commentID)
			assert.Nil(t, err)
			assert.NotZero(t, id)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_Update(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		user     model.User
		purchase model.Purchase
		comment  model.Comment
		update   model.Comment
	}{
		{
			name: "all ok",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
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
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			tc.comment.UserID = userID
			tc.comment.PurchaseID = purchaseID
			commentID, err := repos.Comment.Create(tc.comment)
			assert.Nil(t, err)

			tc.update.UserID = userID
			tc.update.PurchaseID = purchaseID
			id, err := repos.Comment.Update(commentID, tc.update)
			assert.Nil(t, err)
			assert.NotZero(t, id)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindById(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
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
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			if tc.isOk {
				tc.comment.UserID = userID
				tc.comment.PurchaseID = purchaseID
				id, err = repos.Comment.Create(tc.comment)
				assert.Nil(t, err)
				tc.exp.UserID = userID
				tc.exp.PurchaseID = purchaseID
			}
			p, err := repos.Comment.FindByID(id)
			assert.Nil(t, err)
			tc.exp.Date = p.Date
			tc.exp.ID = p.ID
			assert.Equal(t, tc.exp, p)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindAllByUserID(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
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
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(t, err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindAllByUserID(userID)
			assert.Nil(t, err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(t, tc.exp, c)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindByPurchaseID(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
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
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(t, err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindByPurchaseID(purchaseID)
			assert.Nil(t, err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(t, tc.exp, c)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindByUserIDAndPurchaseID(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
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
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(t, err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindByUserIDAndPurchaseID(userID, purchaseID)
			assert.Nil(t, err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(t, tc.exp, c)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindAll(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
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
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(t, err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindAll()
			assert.Nil(t, err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(t, tc.exp, c)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindByText(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
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
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(t, err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindByText(tc.text)
			assert.Nil(t, err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(t, tc.exp, c)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestCommentRepo_FindByPeriod(t *testing.T) {

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk     bool
		name     string
		user     model.User
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
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
			purchase: model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name",
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
			_, err := db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(t, err)

			tc.purchase.UserID = userID
			purchaseID, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(t, err)

			if tc.isOk {
				for i := range tc.comments {
					tc.comments[i].UserID = userID
					tc.comments[i].PurchaseID = purchaseID
					_, err = repos.Comment.Create(tc.comments[i])
					assert.Nil(t, err)
				}
				for i := range tc.exp {
					tc.exp[i].UserID = userID
					tc.exp[i].PurchaseID = purchaseID
				}

			}
			c, err := repos.Comment.FindByPeriod(tc.start, tc.end)
			assert.Nil(t, err)
			for i := range c {
				tc.exp[i].Date = c[i].Date
				tc.exp[i].ID = c[i].ID
			}

			assert.Equal(t, tc.exp, c)

			_, err = db.Exec("DELETE FROM comment")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(t, err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(t, err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}
