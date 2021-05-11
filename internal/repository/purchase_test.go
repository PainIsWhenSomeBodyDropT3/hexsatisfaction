package repository

import (
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	"github.com/stretchr/testify/require"
)

func TestPurchaseRepo_Create(t *testing.T) {
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
				UserId:   1,
				Date:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				FileName: "some name",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := db.Exec("DELETE FROM purchase")
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			tc.purchase.UserId = userId
			id, err := repos.Purchase.Create(tc.purchase)
			require.NoError(t, err)
			require.NotZero(t, id)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_Delete(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			tc.purchase.UserId = userId
			id, err := repos.Purchase.Create(tc.purchase)
			require.NoError(t, err)

			delId, err := repos.Purchase.Delete(id)
			require.NoError(t, err)
			require.NotZero(t, delId)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindById(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				tc.purchase.UserId = userId
				id, err = repos.Purchase.Create(tc.purchase)
				require.NoError(t, err)
				tc.expPurchase.UserId = userId
			}
			p, err := repos.Purchase.FindById(id)
			require.NoError(t, err)
			tc.expPurchase.Date = p.Date
			tc.expPurchase.ID = p.ID
			require.Equal(t, tc.expPurchase, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindLastByUserId(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				tc.expPurchase.UserId = userId
			}
			p, err := repos.Purchase.FindLastByUserId(userId)
			require.NoError(t, err)
			tc.expPurchase.Date = p.Date
			tc.expPurchase.ID = p.ID

			require.Equal(t, tc.expPurchase, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindAllByUserId(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindAllByUserId(userId)
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindByUserIdAndPeriod(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindByUserIdAndPeriod(userId, tc.start, tc.end)
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindByUserIdAfterDate(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindByUserIdAfterDate(userId, tc.start)
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindByUserIdBeforeDate(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindByUserIdBeforeDate(userId, tc.end)
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindByUserIdAndFileName(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindByUserIdAndFileName(userId, tc.fileName)
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindLast(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				tc.expPurchase.UserId = userId
			}
			p, err := repos.Purchase.FindLast()
			require.NoError(t, err)
			tc.expPurchase.Date = p.Date
			tc.expPurchase.ID = p.ID
			require.Equal(t, tc.expPurchase, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindAll(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindAll()
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindByPeriod(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindByPeriod(tc.start, tc.end)
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindAfterDate(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindAfterDate(tc.start)
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindBeforeDate(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindBeforeDate(tc.end)
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}

func TestPurchaseRepo_FindFileName(t *testing.T) {
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
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)

			userId, err := repos.User.Create(tc.user)
			require.NoError(t, err)

			if tc.isOk {
				for i := range tc.purchases {
					tc.purchases[i].UserId = userId
					_, err = repos.Purchase.Create(tc.purchases[i])
					require.NoError(t, err)
				}
				for i := range tc.expPurchases {
					tc.expPurchases[i].UserId = userId
				}
			}
			p, err := repos.Purchase.FindByFileName(tc.fileName)
			require.NoError(t, err)
			for i := range p {
				tc.expPurchases[i].Date = p[i].Date
				tc.expPurchases[i].ID = p[i].ID
			}
			require.Equal(t, tc.expPurchases, p)

			_, err = db.Exec("DELETE FROM purchase")
			require.NoError(t, err)

			_, err = db.Exec("DELETE FROM users")
			require.NoError(t, err)
		})
	}
}
