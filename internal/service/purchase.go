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
		UserId:   request.UserId,
		Date:     request.Date,
		FileName: request.FileName,
	}
	id, err := p.Purchase.Create(purchase)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't create purchase")
	}

	return id, nil
}

// Delete deletes Purchase and returns deleted id.
func (p PurchaseService) Delete(request model.DeletePurchaseRequest) (int, error) {
	id, err := p.Purchase.Delete(request.Id)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't delete purchase")
	}

	return id, nil
}

// FindById finds Purchase by id.
func (p PurchaseService) FindById(request model.FindByIdPurchaseRequest) (*model.Purchase, error) {
	purchase, err := p.Purchase.FindById(request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchase")
	}

	return purchase, nil
}

// FindLastByUserId finds last Purchase by userId.
func (p PurchaseService) FindLastByUserId(request model.FindLastByUserIdPurchaseRequest) (*model.Purchase, error) {
	purchase, err := p.Purchase.FindLastByUserId(request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchase")
	}

	return purchase, nil
}

// FindAllByUserId finds all Purchase by userId.
func (p PurchaseService) FindAllByUserId(request model.FindAllByUserIdPurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindAllByUserId(request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByUserIdAndPeriod finds all Purchase by userId and date period.
func (p PurchaseService) FindByUserIdAndPeriod(request model.FindByUserIdAndPeriodPurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByUserIdAndPeriod(request.Id, request.Start, request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByUserIdAfterDate finds all Purchase by userId and after date.
func (p PurchaseService) FindByUserIdAfterDate(request model.FindByUserIdAfterDatePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByUserIdAfterDate(request.Id, request.Start)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByUserIdBeforeDate finds all Purchase by userId and before date.
func (p PurchaseService) FindByUserIdBeforeDate(request model.FindByUserIdBeforeDatePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByUserIdBeforeDate(request.Id, request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByUserIdAndFileName finds all Purchase by userId and file name.
func (p PurchaseService) FindByUserIdAndFileName(request model.FindByUserIdAndFileNamePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByUserIdAndFileName(request.Id, request.FileName)
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
func (p PurchaseService) FindByPeriod(request model.FindByPeriodPurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByPeriod(request.Start, request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindAfterDate finds all Purchase after date.
func (p PurchaseService) FindAfterDate(request model.FindAfterDatePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindAfterDate(request.Start)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindBeforeDate finds all Purchase before date.
func (p PurchaseService) FindBeforeDate(request model.FindBeforeDatePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindBeforeDate(request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}

// FindByFileName finds all Purchase by file name.
func (p PurchaseService) FindByFileName(request model.FindByFileNamePurchaseRequest) ([]model.Purchase, error) {
	purchases, err := p.Purchase.FindByFileName(request.FileName)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find purchases")
	}

	return purchases, nil
}
