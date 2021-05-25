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

func deleteFileData(assertions *testAssert.Assertions, db *sql.DB) {
	_, err := db.Exec("DELETE FROM file")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM author")
	assertions.Nil(err)
	_, err = db.Exec("DELETE FROM users")
	assertions.Nil(err)
}

func TestFileRepo_Create(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		user   model.User
		author model.Author
		file   model.File
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
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)

			tc.file.AuthorID = authorID
			id, err := repos.File.Create(tc.file)
			assert.Nil(err)
			assert.NotZero(id)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_Update(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		file   model.File
		update model.File
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
			update: model.File{
				Name:        "update",
				Description: "update",
				Size:        1,
				Path:        "update",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      true,
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
			update: model.File{
				Name:        "update",
				Description: "update",
				Size:        1,
				Path:        "update",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      true,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var fileID int
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				fileID, err = repos.File.Create(tc.file)
				assert.Nil(err)
			}

			id, err := repos.File.Update(fileID, tc.file)
			assert.Nil(err)
			assert.Equal(fileID, id)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_Delete(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		file   model.File
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
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var fileID int
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				fileID, err = repos.File.Create(tc.file)
				assert.Nil(err)
			}

			id, err := repos.File.Delete(fileID)
			assert.Nil(err)
			assert.Equal(fileID, id)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_DeleteByAuthorID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		file   model.File
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
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var authorID int
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			if tc.isOk {
				tc.author.UserID = userID
				authorID, err = repos.Author.Create(tc.author)
				assert.Nil(err)

				tc.file.AuthorID = authorID
				_, err = repos.File.Create(tc.file)
				assert.Nil(err)
			}

			id, err := repos.File.DeleteByAuthorID(authorID)
			assert.Nil(err)
			assert.Equal(authorID, id)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_FindByID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		file   model.File
		exp    *model.File
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
			exp: &model.File{},
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
			exp: &model.File{
				Name:        "test",
				Description: "test",
				Size:        1,
				Path:        "test",
				AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Actual:      false,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var fileID int
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				fileID, err = repos.File.Create(tc.file)
				assert.Nil(err)
				tc.exp.AuthorID = authorID
			}

			file, err := repos.File.FindByID(fileID)
			assert.Nil(err)
			tc.exp.ID = file.ID
			tc.exp.AddDate = file.AddDate
			tc.exp.UpdateDate = file.UpdateDate
			assert.Equal(tc.exp, file)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_FindByName(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		file   model.File
		exp    []model.File
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
			exp: []model.File{
				{
					Name:        "test",
					Description: "test",
					Size:        1,
					Path:        "test",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      false,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				_, err := repos.File.Create(tc.file)
				assert.Nil(err)
			}

			files, err := repos.File.FindByName(tc.file.Name)
			assert.Nil(err)
			for i := range files {
				tc.exp[i].ID = files[i].ID
				tc.exp[i].AddDate = files[i].AddDate
				tc.exp[i].UpdateDate = files[i].UpdateDate
				tc.exp[i].AuthorID = authorID
			}
			assert.Equal(tc.exp, files)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_FindAll(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		file   model.File
		exp    []model.File
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
			exp: []model.File{
				{
					Name:        "test",
					Description: "test",
					Size:        1,
					Path:        "test",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      false,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				_, err := repos.File.Create(tc.file)
				assert.Nil(err)
			}

			files, err := repos.File.FindAll()
			assert.Nil(err)
			for i := range files {
				tc.exp[i].ID = files[i].ID
				tc.exp[i].AddDate = files[i].AddDate
				tc.exp[i].UpdateDate = files[i].UpdateDate
				tc.exp[i].AuthorID = authorID
			}
			assert.Equal(tc.exp, files)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_FindByAuthorID(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		file   model.File
		exp    []model.File
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
			exp: []model.File{
				{
					Name:        "test",
					Description: "test",
					Size:        1,
					Path:        "test",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      false,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				_, err := repos.File.Create(tc.file)
				assert.Nil(err)
			}

			files, err := repos.File.FindByAuthorID(authorID)
			assert.Nil(err)
			for i := range files {
				tc.exp[i].ID = files[i].ID
				tc.exp[i].AddDate = files[i].AddDate
				tc.exp[i].UpdateDate = files[i].UpdateDate
				tc.exp[i].AuthorID = authorID
			}
			assert.Equal(tc.exp, files)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_FindNotActual(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		file   model.File
		exp    []model.File
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
			exp: []model.File{
				{
					Name:        "test",
					Description: "test",
					Size:        1,
					Path:        "test",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      false,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				_, err := repos.File.Create(tc.file)
				assert.Nil(err)
			}

			files, err := repos.File.FindNotActual()
			assert.Nil(err)
			for i := range files {
				tc.exp[i].ID = files[i].ID
				tc.exp[i].AddDate = files[i].AddDate
				tc.exp[i].UpdateDate = files[i].UpdateDate
				tc.exp[i].AuthorID = authorID
			}
			assert.Equal(tc.exp, files)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_FindActual(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		user   model.User
		author model.Author
		file   model.File
		exp    []model.File
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
				Actual:      true,
			},
			exp: []model.File{
				{
					Name:        "test",
					Description: "test",
					Size:        1,
					Path:        "test",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      true,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				_, err := repos.File.Create(tc.file)
				assert.Nil(err)
			}

			files, err := repos.File.FindActual()
			assert.Nil(err)
			for i := range files {
				tc.exp[i].ID = files[i].ID
				tc.exp[i].AddDate = files[i].AddDate
				tc.exp[i].UpdateDate = files[i].UpdateDate
				tc.exp[i].AuthorID = authorID
			}
			assert.Equal(tc.exp, files)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_FindAddedByPeriod(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		start  time.Time
		end    time.Time
		user   model.User
		author model.Author
		file   model.File
		exp    []model.File
	}{
		{
			name:  "not found",
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
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
			name:  "all ok",
			isOk:  true,
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
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
			exp: []model.File{
				{
					Name:        "test",
					Description: "test",
					Size:        1,
					Path:        "test",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      false,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				_, err := repos.File.Create(tc.file)
				assert.Nil(err)
			}

			files, err := repos.File.FindAddedByPeriod(tc.start, tc.end)
			assert.Nil(err)
			for i := range files {
				tc.exp[i].ID = files[i].ID
				tc.exp[i].AddDate = files[i].AddDate
				tc.exp[i].UpdateDate = files[i].UpdateDate
				tc.exp[i].AuthorID = authorID
			}
			assert.Equal(tc.exp, files)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}

func TestFileRepo_FindUpdatedByPeriod(t *testing.T) {
	assert := testAssert.New(t)
	db, repos, err := Connect2Repositories()
	require.NoError(t, err)
	tt := []struct {
		name   string
		isOk   bool
		start  time.Time
		end    time.Time
		user   model.User
		author model.Author
		file   model.File
		exp    []model.File
	}{
		{
			name:  "not found",
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
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
			name:  "all ok",
			isOk:  true,
			start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			end:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC),
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
			exp: []model.File{
				{
					Name:        "test",
					Description: "test",
					Size:        1,
					Path:        "test",
					AddDate:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UpdateDate:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					Actual:      false,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deleteFileData(assert, db)

			userID, err := repos.User.Create(tc.user)
			assert.Nil(err)
			tc.author.UserID = userID
			authorID, err := repos.Author.Create(tc.author)
			assert.Nil(err)
			if tc.isOk {
				tc.file.AuthorID = authorID
				_, err := repos.File.Create(tc.file)
				assert.Nil(err)
			}

			files, err := repos.File.FindUpdatedByPeriod(tc.start, tc.end)
			assert.Nil(err)
			for i := range files {
				tc.exp[i].ID = files[i].ID
				tc.exp[i].AddDate = files[i].AddDate
				tc.exp[i].UpdateDate = files[i].UpdateDate
				tc.exp[i].AuthorID = authorID
			}
			assert.Equal(tc.exp, files)

			deleteFileData(assert, db)
		})
	}
	err = db.Close()
	require.NoError(t, err)
}
