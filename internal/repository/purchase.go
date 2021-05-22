package repository

import (
	"database/sql"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
)

// PurchaseRepo is a purchase repository.
type PurchaseRepo struct {
	db *sql.DB
}

// NewPurchaseRepo is a PurchaseRepo constructor.
func NewPurchaseRepo(db *sql.DB) *PurchaseRepo {
	return &PurchaseRepo{db: db}
}

// Create creates new purchase and returns id.
func (p PurchaseRepo) Create(purchase model.Purchase) (int, error) {
	var creatID int
	rows, err := p.db.Query("INSERT INTO purchase (userID , date,fileID) VALUES ($1,$2,$3) RETURNING id ", purchase.UserID, purchase.Date, purchase.FileID)
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

// Delete deletes purchase and returns deleted id.
func (p PurchaseRepo) Delete(id int) (int, error) {
	var delID int
	rows, err := p.db.Query("DELETE FROM purchase WHERE id=$1 RETURNING id ", id)
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

// FindByID finds purchase by id.
func (p PurchaseRepo) FindByID(id int) (*model.Purchase, error) {
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
	}

	return &purchase, rows.Err()
}

// FindLastByUserID finds last purchase by user id.
func (p PurchaseRepo) FindLastByUserID(id int) (*model.Purchase, error) {
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID = $1 ORDER BY id DESC limit 1;", id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
	}

	return &purchase, rows.Err()
}

// FindAllByUserID finds purchases by user id.
func (p PurchaseRepo) FindAllByUserID(id int) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1", id)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByUserIDAndPeriod finds purchases by user id and date period.
func (p PurchaseRepo) FindByUserIDAndPeriod(id int, start, end time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1 AND date BETWEEN $2 AND $3", id, start, end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByUserIDAfterDate finds purchases by user id and after date.
func (p PurchaseRepo) FindByUserIDAfterDate(id int, start time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1 AND date >= $2", id, start)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByUserIDBeforeDate finds purchases by user id and before date.
func (p PurchaseRepo) FindByUserIDBeforeDate(id int, end time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1 AND date <= $2", id, end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByUserIDAndFileID finds purchases by user id and file id.
func (p PurchaseRepo) FindByUserIDAndFileID(userID, fileID int) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1 AND fileID = $2", userID, fileID)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindLast finds last purchase.
func (p PurchaseRepo) FindLast() (*model.Purchase, error) {
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase ORDER BY id DESC limit 1;")
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
	}

	return &purchase, rows.Err()
}

// FindAll finds purchases.
func (p PurchaseRepo) FindAll() ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase")
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByPeriod finds purchases by date period.
func (p PurchaseRepo) FindByPeriod(start, end time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE date BETWEEN $1 AND $2", start, end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindAfterDate finds purchases after date.
func (p PurchaseRepo) FindAfterDate(start time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE  date >= $1", start)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindBeforeDate finds purchases before date.
func (p PurchaseRepo) FindBeforeDate(end time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE date <= $1", end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByFileID finds purchases by file id.
func (p PurchaseRepo) FindByFileID(id int) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE  fileID = $1", id)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserID, &purchase.Date, &purchase.FileID)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}
