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

// Create creates new Purchase and returns id.
func (p PurchaseRepo) Create(purchase model.Purchase) (int, error) {
	var creatId int
	rows, err := p.db.Query("INSERT INTO purchase (userID , date,fileName) VALUES ($1,$2,$3) RETURNING id ", purchase.UserId, purchase.Date, purchase.FileName)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&creatId)
		if err != nil {
			return 0, err
		}
	}
	return creatId, rows.Err()

}

// Delete deletes Purchase and returns deleted id.
func (p PurchaseRepo) Delete(id int) (int, error) {
	var delId int
	rows, err := p.db.Query("DELETE FROM purchase WHERE id=$1 RETURNING id ", id)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&delId)
		if err != nil {
			return 0, err
		}
	}

	return delId, rows.Err()
}

// FindById finds Purchase by id.
func (p PurchaseRepo) FindById(id int) (*model.Purchase, error) {
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
	}

	return &purchase, rows.Err()
}

// FindLastByUserId finds last Purchase by userId.
func (p PurchaseRepo) FindLastByUserId(id int) (*model.Purchase, error) {
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID = $1 ORDER BY id DESC limit 1;", id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
	}

	return &purchase, rows.Err()
}

// FindAllByUserId finds all Purchase by userId.
func (p PurchaseRepo) FindAllByUserId(id int) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1", id)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByUserIdAndPeriod finds all Purchase by userId and date period.
func (p PurchaseRepo) FindByUserIdAndPeriod(id int, start, end time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1 AND date BETWEEN $2 AND $3", id, start, end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByUserIdAfterDate finds all Purchase by userId and after date.
func (p PurchaseRepo) FindByUserIdAfterDate(id int, start time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1 AND date >= $2", id, start)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByUserIdBeforeDate finds all Purchase by userId and before date.
func (p PurchaseRepo) FindByUserIdBeforeDate(id int, end time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1 AND date <= $2", id, end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByUserIdAndFileName finds all Purchase by userId and file name.
func (p PurchaseRepo) FindByUserIdAndFileName(id int, name string) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE userID=$1 AND fileName = $2", id, name)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindLast finds last Purchase.
func (p PurchaseRepo) FindLast() (*model.Purchase, error) {
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase ORDER BY id DESC limit 1;")
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
	}

	return &purchase, rows.Err()
}

// FindAll finds all Purchase.
func (p PurchaseRepo) FindAll() ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase")
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByPeriod finds all Purchase by date period.
func (p PurchaseRepo) FindByPeriod(start, end time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE date BETWEEN $1 AND $2", start, end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindAfterDate finds all Purchase after date.
func (p PurchaseRepo) FindAfterDate(start time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE  date >= $1", start)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindBeforeDate finds all Purchase before date.
func (p PurchaseRepo) FindBeforeDate(end time.Time) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE date <= $1", end)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}

// FindByFileName finds all Purchase by file name.
func (p PurchaseRepo) FindByFileName(name string) ([]model.Purchase, error) {
	var purchases []model.Purchase
	var purchase model.Purchase
	rows, err := p.db.Query("SELECT * FROM purchase WHERE  fileName = $1", name)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {

		err = rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Date, &purchase.FileName)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, rows.Err()
}
