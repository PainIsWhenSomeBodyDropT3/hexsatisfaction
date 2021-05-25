package repository

import (
	"database/sql"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
)

// FileRepo is a file repository.
type FileRepo struct {
	db *sql.DB
}

// NewFileRepo is a FileRepo constructor.
func NewFileRepo(db *sql.DB) *FileRepo {
	return &FileRepo{db: db}
}

// Create creates new file and returns id.
func (f FileRepo) Create(file model.File) (int, error) {
	var creatID int
	rows, err := f.db.Query("INSERT INTO file (name, description, size, path, addDate, updateDate, actual, authorID) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id",
		file.Name, file.Description, file.Size, file.Path, file.AddDate, file.UpdateDate, file.Actual, file.AuthorID)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&creatID)
		if err != nil {
			return 0, err
		}
	}
	return creatID, rows.Err()

}

// Update updates file and returns id.
func (f FileRepo) Update(id int, file model.File) (int, error) {
	var updatedID int
	rows, err := f.db.Query("UPDATE file SET name=$1, description=$2, size=$3, path=$4, addDate=$5, updateDate=$6, actual=$7, authorID=$8 WHERE id=$9 RETURNING id",
		file.Name, file.Description, file.Size, file.Path, file.AddDate, file.UpdateDate, file.Actual, file.AuthorID, id)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&updatedID)
		if err != nil {
			return 0, err
		}
	}
	return updatedID, rows.Err()
}

// Delete deletes file and returns deleted id.
func (f FileRepo) Delete(id int) (int, error) {
	var delID int
	rows, err := f.db.Query("DELETE FROM file WHERE id=$1 RETURNING id ", id)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&delID)
		if err != nil {
			return 0, err
		}
	}

	return delID, rows.Err()
}

// DeleteByAuthorID deletes file by authorID and returns authorID.
func (f FileRepo) DeleteByAuthorID(id int) (int, error) {
	var delID int
	rows, err := f.db.Query("DELETE FROM file WHERE authorID=$1 RETURNING authorID ", id)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&delID)
		if err != nil {
			return 0, err
		}
	}

	return delID, rows.Err()
}

// FindByID finds file by id.
func (f FileRepo) FindByID(id int) (*model.File, error) {
	var file model.File
	rows, err := f.db.Query("SELECT * FROM file WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&file.ID, &file.Name, &file.Description, &file.Size, &file.Path, &file.AddDate, &file.UpdateDate, &file.Actual, &file.AuthorID)
		if err != nil {
			return nil, err
		}
	}

	return &file, rows.Err()
}

// FindByName finds files by name.
func (f FileRepo) FindByName(name string) ([]model.File, error) {
	var files []model.File
	var file model.File
	rows, err := f.db.Query("SELECT * FROM file WHERE name=$1", name)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&file.ID, &file.Name, &file.Description, &file.Size, &file.Path, &file.AddDate, &file.UpdateDate, &file.Actual, &file.AuthorID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, rows.Err()
}

// FindAll finds files.
func (f FileRepo) FindAll() ([]model.File, error) {
	var files []model.File
	var file model.File
	rows, err := f.db.Query("SELECT * FROM file")
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&file.ID, &file.Name, &file.Description, &file.Size, &file.Path, &file.AddDate, &file.UpdateDate, &file.Actual, &file.AuthorID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, rows.Err()
}

// FindByAuthorID finds files by author id.
func (f FileRepo) FindByAuthorID(id int) ([]model.File, error) {
	var files []model.File
	var file model.File
	rows, err := f.db.Query("SELECT * FROM file WHERE authorID=$1", id)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&file.ID, &file.Name, &file.Description, &file.Size, &file.Path, &file.AddDate, &file.UpdateDate, &file.Actual, &file.AuthorID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, rows.Err()
}

// FindNotActual finds not actual files.
func (f FileRepo) FindNotActual() ([]model.File, error) {
	var files []model.File
	var file model.File
	rows, err := f.db.Query("SELECT * FROM file WHERE actual=$1", false)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&file.ID, &file.Name, &file.Description, &file.Size, &file.Path, &file.AddDate, &file.UpdateDate, &file.Actual, &file.AuthorID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, rows.Err()
}

// FindActual finds actual files.
func (f FileRepo) FindActual() ([]model.File, error) {
	var files []model.File
	var file model.File
	rows, err := f.db.Query("SELECT * FROM file WHERE actual=$1", true)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&file.ID, &file.Name, &file.Description, &file.Size, &file.Path, &file.AddDate, &file.UpdateDate, &file.Actual, &file.AuthorID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, rows.Err()
}

// FindAddedByPeriod finds added files by date period.
func (f FileRepo) FindAddedByPeriod(start, end time.Time) ([]model.File, error) {
	var files []model.File
	var file model.File
	rows, err := f.db.Query("SELECT * FROM file WHERE addDate BETWEEN $1 AND $2", start, end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&file.ID, &file.Name, &file.Description, &file.Size, &file.Path, &file.AddDate, &file.UpdateDate, &file.Actual, &file.AuthorID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, rows.Err()
}

// FindUpdatedByPeriod finds updated files by date period.
func (f FileRepo) FindUpdatedByPeriod(start, end time.Time) ([]model.File, error) {
	var files []model.File
	var file model.File
	rows, err := f.db.Query("SELECT * FROM file WHERE updateDate BETWEEN $1 AND $2", start, end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&file.ID, &file.Name, &file.Description, &file.Size, &file.Path, &file.AddDate, &file.UpdateDate, &file.Actual, &file.AuthorID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, rows.Err()
}
