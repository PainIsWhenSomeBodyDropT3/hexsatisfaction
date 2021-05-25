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

func deletePurchaseData(assertions *testAssert.Assertions, db *sql.DB) {
	_, err := db.Exec("DELETE FROM purchase")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM file")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM author")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM users")
	assertions.Nil(err)
}

func TestPurchaseRepo_Create(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		user     model.User
		author   model.Author
		file     model.File
		purchase model.Purchase
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
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)
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
			id, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)
			assert.NotZero(id)
			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_Delete(t *testing.T) {
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
		},
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
				Date: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var purchaseID int
			deletePurchaseData(assert, db)

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
				purchaseID, err = repos.Purchase.Create(tc.purchase)
				assert.Nil(err)
			}

			delID, err := repos.Purchase.Delete(purchaseID)
			assert.Nil(err)
			assert.Equal(purchaseID, delID)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_DeleteByFileID(t *testing.T) {
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
		},
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
				Date: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var fileID int
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			if tc.isOk {
				tc.file.AuthorID = authorID
				fileID, err = repos.File.Create(tc.file)
				assert.Nil(err)

				tc.purchase.UserID = userID
				tc.purchase.FileID = fileID
				_, err = repos.Purchase.Create(tc.purchase)
				assert.Nil(err)
			}

			delID, err := repos.Purchase.DeleteByFileID(fileID)
			assert.Nil(err)
			assert.Equal(fileID, delID)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindById(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk        bool
		name        string
		user        model.User
		author      model.Author
		file        model.File
		purchase    model.Purchase
		expPurchase *model.Purchase
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
			expPurchase: &model.Purchase{},
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
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
			expPurchase: &model.Purchase{
				Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var id int
			deletePurchaseData(assert, db)

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
				id, err = repos.Purchase.Create(tc.purchase)
				assert.Nil(err)
				tc.expPurchase.UserID = userID
				tc.expPurchase.FileID = fileID
			}
			p, err := repos.Purchase.FindByID(id)
			assert.Nil(err)
			tc.expPurchase.Date = p.Date
			tc.expPurchase.ID = p.ID
			assert.Equal(tc.expPurchase, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindLastByUserId(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk        bool
		name        string
		user        model.User
		author      model.Author
		file        model.File
		purchases   []model.Purchase
		expPurchase *model.Purchase
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
			expPurchase: &model.Purchase{},
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
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchase: &model.Purchase{
				Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				tc.expPurchase.UserID = userID
				tc.expPurchase.FileID = fileID
			}
			p, err := repos.Purchase.FindLastByUserID(userID)
			assert.Nil(err)
			tc.expPurchase.Date = p.Date
			tc.expPurchase.ID = p.ID

			assert.Equal(tc.expPurchase, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindAllByUserId(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		purchases    []model.Purchase
		expPurchases []model.Purchase
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
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindAllByUserID(userID)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindByUserIdAndPeriod(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		start        time.Time
		end          time.Time
		purchases    []model.Purchase
		expPurchases []model.Purchase
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
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
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
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindByUserIDAndPeriod(userID, tc.start, tc.end)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindByUserIdAfterDate(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		start        time.Time
		purchases    []model.Purchase
		expPurchases []model.Purchase
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
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
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
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindByUserIDAfterDate(userID, tc.start)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindByUserIdBeforeDate(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		end          time.Time
		purchases    []model.Purchase
		expPurchases []model.Purchase
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
			end: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
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
			end: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindByUserIDBeforeDate(userID, tc.end)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindByUserIdAndFileID(t *testing.T) {
	assert := testAssert.New(t)

	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		purchases    []model.Purchase
		expPurchases []model.Purchase
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
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindByUserIDAndFileID(userID, fileID)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindLast(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk        bool
		name        string
		user        model.User
		author      model.Author
		file        model.File
		purchases   []model.Purchase
		expPurchase *model.Purchase
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
			expPurchase: &model.Purchase{},
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
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchase: &model.Purchase{
				Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				tc.expPurchase.UserID = userID
				tc.expPurchase.FileID = fileID
			}
			p, err := repos.Purchase.FindLast()
			assert.Nil(err)
			tc.expPurchase.Date = p.Date
			tc.expPurchase.ID = p.ID
			assert.Equal(tc.expPurchase, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindAll(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		purchases    []model.Purchase
		expPurchases []model.Purchase
	}{
		{
			name: "find err",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
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
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindAll()
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindByPeriod(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		start        time.Time
		end          time.Time
		purchases    []model.Purchase
		expPurchases []model.Purchase
	}{
		{
			name: "find err",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
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
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindByPeriod(tc.start, tc.end)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindAfterDate(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		start        time.Time
		purchases    []model.Purchase
		expPurchases []model.Purchase
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
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindAfterDate(tc.start)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindBeforeDate(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		end          time.Time
		purchases    []model.Purchase
		expPurchases []model.Purchase
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
			end: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			end: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindBeforeDate(tc.end)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindFileID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		author       model.Author
		file         model.File
		purchases    []model.Purchase
		expPurchases []model.Purchase
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
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},

			purchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
			expPurchases: []model.Purchase{
				{
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				},
				{
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deletePurchaseData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			fileID, err := repos.File.Create(tc.file)
			assert.Nil(err)
			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					tc.purchases[i].FileID = fileID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
					tc.expPurchases[i].FileID = fileID
				}
			}
			p, err := repos.Purchase.FindByFileID(fileID)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			deletePurchaseData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}
