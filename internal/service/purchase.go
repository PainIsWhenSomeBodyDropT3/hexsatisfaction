package service

import (
	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/pkg/errors"
)

// PurchaseService is a purchase service.
type PurchaseService struct {
	repository.Purchase
}

// NewPurchaseService is a PurchaseService service constructor.
func NewPurchaseService(purchase repository.Purchase) *PurchaseService {
	return &PurchaseService{purchase}
}

// Create creates new Purchase and returns id.
func (p PurchaseService) Create(request model.CreatePurchaseRequest) (int, error) {
	purchase := model.Purchase{
		UserID: request.UserID,
		Date:   request.Date,
		FileID: request.FileID,
	}
	id, err := p.Purchase.Create(purchase)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't create purchase")
	}

	return id, nil
}

// Delete deletes Purchase and returns deleted id.
func (p PurchaseService) Delete(request model.DeletePurchaseRequest) (int, error) {
	id, err := p.Purchase.Delete(request.ID)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't delete purchase")
	}

	return id, nil
}

// FindByID finds Purchase by id.
func (p PurchaseService) FindByID(request model.IDPurchaseRequest) (*model.Purchase, error) {
	purchase, err := p.Purchase.FindByID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchase")
	}

	return purchase, nil
}

// FindLastByUserID finds last Purchase by userId.
func (p PurchaseService) FindLastByUserID(request model.UserIDPurchaseRequest) (*model.Purchase, error) {
	purchase, err := p.Purchase.FindLastByUserID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchase")
	}

	return purchase, nil
}

// FindAllByUserID finds all Purchase by userId.
func (p PurchaseService) FindAllByUserID(request model.UserIDPurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindAllByUserID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByUserIDAndPeriod finds all Purchase by userId and date period.
func (p PurchaseService) FindByUserIDAndPeriod(request model.UserIDPeriodPurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByUserIDAndPeriod(request.ID, request.Start, request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByUserIDAfterDate finds all Purchase by userId and after date.
func (p PurchaseService) FindByUserIDAfterDate(request model.UserIDAfterDatePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByUserIDAfterDate(request.ID, request.Start)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByUserIDBeforeDate finds all Purchase by userId and before date.
func (p PurchaseService) FindByUserIDBeforeDate(request model.UserIDBeforeDatePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByUserIDBeforeDate(request.ID, request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByUserIDAndFileID finds all Purchase by userId and file name.
func (p PurchaseService) FindByUserIDAndFileID(request model.UserIDFileIDPurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByUserIDAndFileID(request.UserID, request.FileID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindLast finds last Purchase.
func (p PurchaseService) FindLast() (*model.Purchase, error) {
	purchase, err := p.Purchase.FindLast()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchase")
	}

	return purchase, nil
}

// FindAll finds all Purchase.
func (p PurchaseService) FindAll() ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByPeriod finds all Purchase by date period.
func (p PurchaseService) FindByPeriod(request model.PeriodPurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByPeriod(request.Start, request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindAfterDate finds all Purchase after date.
func (p PurchaseService) FindAfterDate(request model.AfterDatePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindAfterDate(request.Start)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindBeforeDate finds all Purchase before date.
func (p PurchaseService) FindBeforeDate(request model.BeforeDatePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindBeforeDate(request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByFileID finds all Purchase by file name.
func (p PurchaseService) FindByFileID(request model.FileIDPurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByFileID(request.FileID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}
