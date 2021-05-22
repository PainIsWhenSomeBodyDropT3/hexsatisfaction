package repository

import (
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	testAssert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPurchaseRepo_Create(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name     string
		user     model.User
		purchase model.Purchase
	}{
		{
			name: "all ok",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			purchase: model.Purchase{
				UserID:   1,
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			tc.purchase.UserID = userID
			id, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)
			assert.NotZero(id)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
		user     model.User
		purchase model.Purchase
	}{
		{
			name: "all ok",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			purchase: model.Purchase{
				Date:     time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
				FileName: "some name",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			tc.purchase.UserID = userID
			id, err := repos.Purchase.Create(tc.purchase)
			assert.Nil(err)

			delID, err := repos.Purchase.Delete(id)
			assert.Nil(err)
			assert.NotZero(delID)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name1",
			},
			expPurchase: &model.Purchase{
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name1",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var id int
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				tc.purchase.UserID = userID
				id, err = repos.Purchase.Create(tc.purchase)
				assert.Nil(err)
				tc.expPurchase.UserID = userID
			}
			p, err := repos.Purchase.FindByID(id)
			assert.Nil(err)
			tc.expPurchase.Date = p.Date
			tc.expPurchase.ID = p.ID
			assert.Equal(tc.expPurchase, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
			purchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchase: &model.Purchase{
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name2",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				tc.expPurchase.UserID = userID
			}
			p, err := repos.Purchase.FindLastByUserID(userID)
			assert.Nil(err)
			tc.expPurchase.Date = p.Date
			tc.expPurchase.ID = p.ID

			assert.Equal(tc.expPurchase, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
			purchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindAllByUserID(userID)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			purchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindByUserIDAndPeriod(userID, tc.start, tc.end)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindByUserIDAfterDate(userID, tc.start)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindByUserIDBeforeDate(userID, tc.end)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindByUserIdAndFileName(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		fileName     string
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
			fileName: "wrong name",
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			fileName: "some name1",
			purchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindByUserIDAndFileName(userID, tc.fileName)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
			purchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchase: &model.Purchase{
				Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				FileName: "some name2",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				tc.expPurchase.UserID = userID
			}
			p, err := repos.Purchase.FindLast()
			assert.Nil(err)
			tc.expPurchase.Date = p.Date
			tc.expPurchase.ID = p.ID
			assert.Equal(tc.expPurchase, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
			purchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindAll()
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			purchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindByPeriod(tc.start, tc.end)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindAfterDate(tc.start)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
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
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindBeforeDate(tc.end)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestPurchaseRepo_FindFileName(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		isOk         bool
		name         string
		user         model.User
		fileName     string
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
			fileName: "wrong name",
		},
		{
			name: "all ok",
			isOk: true,
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			fileName: "some name1",
			purchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
				{
					Date:     time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name2",
				},
			},
			expPurchases: []model.Purchase{
				{
					Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					FileName: "some name1",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			assert.Nil(err)
			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserID = userID
					_, err = repos.Purchase.Create(tc.purchases[i])
					assert.Nil(err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserID = userID
				}
			}
			p, err := repos.Purchase.FindByFileName(tc.fileName)
			assert.Nil(err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			assert.Equal(tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			assert.Nil(err)

			_, err = db.Exec("DELETE FROM users")
			assert.Nil(err)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}
