package service

import (
	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/pkg/errors"
)

// FileService is a file service.
type FileService struct {
	repository.File
	repository.Purchase
	repository.Comment
}

// NewFileService is a FileService service constructor.
func NewFileService(file repository.File, purchase repository.Purchase, comment repository.Comment) *FileService {
	return &FileService{file, purchase, comment}
}

// Create creates new file and returns id.
func (f FileService) Create(request model.CreateFileRequest) (int, error) {
	file := model.File{
		Name:        request.Name,
		Description: request.Description,
		Size:        request.Size,
		Path:        request.Path,
		AddDate:     request.AddDate,
		UpdateDate:  request.UpdateDate,
		Actual:      request.Actual,
		AuthorID:    request.AuthorID,
	}
	id, err := f.File.Create(file)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't create file")
	}

	return id, nil
}

// Update updates file and returns id.
func (f FileService) Update(request model.UpdateFileRequest) (int, error) {
	file := model.File{
		Name:        request.Name,
		Description: request.Description,
		Size:        request.Size,
		Path:        request.Path,
		AddDate:     request.AddDate,
		UpdateDate:  request.UpdateDate,
		Actual:      request.Actual,
		AuthorID:    request.AuthorID,
	}
	id, err := f.File.Update(request.ID, file)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't update file")
	}

	return id, nil
}

// Delete deletes file and returns deleted id.
func (f FileService) Delete(request model.DeleteFileRequest) (int, error) {

	purchases, err := f.Purchase.FindByFileID(request.ID)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't get purchases")
	}

	for _, p := range purchases {
		_, err := f.Comment.DeleteByPurchaseID(p.ID)
		if err != nil {
			return 0, errors.Wrap(err, "couldn't delete comment")
		}
	}

	fileID, err := f.Purchase.DeleteByFileID(request.ID)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't delete purchases")
	}

	id, err := f.File.Delete(fileID)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't delete file")
	}

	return id, nil
}

func (f FileService) removeComments(purchases []model.Purchase) error {
	for _, p := range purchases {
		_, err := f.Comment.DeleteByPurchaseID(p.ID)
		if err != nil {
			return errors.Wrap(err, "couldn't delete comment")
		}
	}

	return nil
}

// FindByID finds file by id.
func (f FileService) FindByID(request model.IDFileRequest) (*model.File, error) {
	file, err := f.File.FindByID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find file")
	}

	return file, nil
}

// FindByName finds files by name.
func (f FileService) FindByName(request model.NameFileRequest) ([]model.File, error) {
	files, err := f.File.FindByName(request.Name)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find files")
	}

	return files, nil
}

// FindAll finds files.
func (f FileService) FindAll() ([]model.File, error) {
	files, err := f.File.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find files")
	}

	return files, nil
}

// FindByAuthorID finds files by author id.
func (f FileService) FindByAuthorID(request model.AuthorIDFileRequest) ([]model.File, error) {
	files, err := f.File.FindByAuthorID(request.ID)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find files")
	}

	return files, nil
}

// FindNotActual finds not actual files.
func (f FileService) FindNotActual() ([]model.File, error) {
	files, err := f.File.FindNotActual()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find files")
	}

	return files, nil
}

// FindActual finds actual files.
func (f FileService) FindActual() ([]model.File, error) {
	files, err := f.File.FindActual()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find files")
	}

	return files, nil
}

// FindAddedByPeriod finds added files by date period.
func (f FileService) FindAddedByPeriod(request model.AddedPeriodFileRequest) ([]model.File, error) {
	files, err := f.File.FindAddedByPeriod(request.Start, request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find files")
	}

	return files, nil
}

// FindUpdatedByPeriod finds updated files by date period.
func (f FileService) FindUpdatedByPeriod(request model.UpdatedPeriodFileRequest) ([]model.File, error) {
	files, err := f.File.FindUpdatedByPeriod(request.Start, request.End)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find files")
	}

	return files, nil
}
