package repository

import (
	"database/sql"
	"testing"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/model/dto"
	testAssert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func deleteAuthorData(assertions *testAssert.Assertions, db *sql.DB) {
	_, err := db.Exec("DELETE FROM author")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM users")
	assertions.Nil(err)
}

func TestAuthorRepo_Create(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		user   model.User
		author model.Author
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
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteAuthorData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			id, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			assert.NotZero(id)

			deleteAuthorData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestAuthorRepo_Update(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		update model.Author
	}{
		{
			name: "not found",
			user: model.User{
				Login:    "test",
				Password: "test",
				RoleID:   dto.USER,
			},
			update: model.Author{
				Name:        "update",
				Age:         1,
				Description: "update",
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
			update: model.Author{
				Name:        "update",
				Age:         1,
				Description: "update",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var authorID int
			deleteAuthorData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)

			if tc.isOk {
				tc.author.UserID = userID
				authorID, err = repos.Author.Create(tc.author)
				assert.Nil(err)
			}

			tc.update.UserID = userID
			id, err := repos.Author.Update(authorID, tc.update)
			assert.Nil(err)
			assert.Equal(authorID, id)

			deleteAuthorData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestAuthorRepo_Delete(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
	}{
		{
			name: "not found",
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
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var authorID int
			deleteAuthorData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			if tc.isOk {
				tc.author.UserID = userID
				authorID, err = repos.Author.Create(tc.author)
				assert.Nil(err)
			}

			id, err := repos.Author.Delete(authorID)
			assert.Nil(err)
			assert.Equal(authorID, id)

			deleteAuthorData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestAuthorRepo_FindByID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		exp    *model.Author
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
			exp: &model.Author{},
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
			exp: &model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var authorID int
			deleteAuthorData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			if tc.isOk {
				tc.author.UserID = userID
				authorID, err = repos.Author.Create(tc.author)
				assert.Nil(err)
				tc.exp.UserID = userID
			}
			author, err := repos.Author.FindByID(authorID)
			assert.Nil(err)
			tc.exp.ID = authorID
			assert.Equal(tc.exp, author)

			deleteAuthorData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestAuthorRepo_FindByUserID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		exp    *model.Author
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
			exp: &model.Author{},
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
			exp: &model.Author{
				Name:        "test",
				Age:         1,
				Description: "test",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var authorID int
			deleteAuthorData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			if tc.isOk {
				tc.author.UserID = userID
				authorID, err = repos.Author.Create(tc.author)
				assert.Nil(err)
				tc.exp.UserID = userID
			}
			author, err := repos.Author.FindByUserID(userID)
			assert.Nil(err)
			tc.exp.ID = authorID
			assert.Equal(tc.exp, author)

			deleteAuthorData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestAuthorRepo_FindByName(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		exp    []model.Author
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
			exp: []model.Author{
				{
					Name:        "test",
					Age:         1,
					Description: "test",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var authorID int
			deleteAuthorData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			if tc.isOk {
				tc.author.UserID = userID
				authorID, err = repos.Author.Create(tc.author)
				assert.Nil(err)
			}
			authors, err := repos.Author.FindByName(tc.author.Name)
			assert.Nil(err)
			for i := range authors {
				tc.exp[i].ID = authorID
				tc.exp[i].UserID = userID
			}
			assert.Equal(tc.exp, authors)

			deleteAuthorData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestAuthorRepo_FindAll(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		exp    []model.Author
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
			exp: []model.Author{
				{
					Name:        "test",
					Age:         1,
					Description: "test",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var authorID int
			deleteAuthorData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			if tc.isOk {
				tc.author.UserID = userID
				authorID, err = repos.Author.Create(tc.author)
				assert.Nil(err)
			}
			authors, err := repos.Author.FindAll()
			assert.Nil(err)
			for i := range authors {
				tc.exp[i].ID = authorID
				tc.exp[i].UserID = userID
			}
			assert.Equal(tc.exp, authors)

			deleteAuthorData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}
