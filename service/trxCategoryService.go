package service

import (
	"errors"
	"myfinnplan/entity"
	"myfinnplan/input"
	"myfinnplan/repository"
	"time"
)

type TrxCategoryService interface {
	CreateTrxCategory(input input.TrxCategoryInput, userName string) (entity.TrxCategory, error)
	EditTrxCategory(input input.TrxCategoryEditInput, userName string) (entity.TrxCategory, error)
	GetTrxCategoryById(id int) ([]entity.TrxCategory, error)
	GetTrxCategoryByCategoryCode(categoryCode string) ([]entity.TrxCategory, error)
	GetTrxCategoryByCategoryName(categoryName string) ([]entity.TrxCategory, error)

	GetAllTrxCategory() ([]entity.TrxCategory, error)
	DeleteTrxCategory(id int, userName string) (entity.TrxCategory, error)
}

type trxCategoryService struct {
	trxCategoryRepository repository.TrxCategoryRepository
}

func NewTrxCategoryService(trxCategoryRepository repository.TrxCategoryRepository) *trxCategoryService {
	return &trxCategoryService{trxCategoryRepository}
}

func (s *trxCategoryService) CreateTrxCategory(input input.TrxCategoryInput, userName string) (entity.TrxCategory, error) {
	trxCategory := entity.TrxCategory{
		CategoryName: input.CategoryName,
		CreatedBy:    userName,
		CreatedDate:  time.Now(),
	}

	newTrxCategory, err := s.trxCategoryRepository.Save(trxCategory)

	if err != nil {
		return trxCategory, err
	}

	return newTrxCategory, nil
}

func (s *trxCategoryService) EditTrxCategory(input input.TrxCategoryEditInput, userName string) (entity.TrxCategory, error) {
	oldTrxCategorys, err := s.trxCategoryRepository.FindById(input.Id)

	if err != nil {
		return entity.TrxCategory{}, err
	}

	oldTrxCategory := oldTrxCategorys[0]

	trxCategory := entity.TrxCategory{
		Id:           input.Id,
		CategoryName: input.CategoryName,
		CreatedBy:    oldTrxCategory.CreatedBy,
		CreatedDate:  oldTrxCategory.CreatedDate,
		UpdatedBy:    userName,
	}

	newTrxCategory, err := s.trxCategoryRepository.Edit(trxCategory)

	if err != nil {
		return trxCategory, err
	}

	return newTrxCategory, nil
}

func (s *trxCategoryService) GetTrxCategoryById(id int) ([]entity.TrxCategory, error) {

	trxCategory, err := s.trxCategoryRepository.FindById(id)

	if err != nil {
		return trxCategory, err
	}

	if len(trxCategory) == 0 {
		return trxCategory, errors.New("trxCategory not found")
	}

	return trxCategory, nil
}
func (s *trxCategoryService) GetTrxCategoryByCategoryCode(categoryCode string) ([]entity.TrxCategory, error) {

	trxCategory, err := s.trxCategoryRepository.FindByCategoryCode(categoryCode)

	if err != nil {
		return trxCategory, err
	}

	if len(trxCategory) == 0 {
		return trxCategory, errors.New("trxCategory not found")
	}

	return trxCategory, nil
}
func (s *trxCategoryService) GetTrxCategoryByCategoryName(categoryName string) ([]entity.TrxCategory, error) {

	trxCategory, err := s.trxCategoryRepository.FindByCategoryName(categoryName)

	if err != nil {
		return trxCategory, err
	}

	if len(trxCategory) == 0 {
		return trxCategory, errors.New("trxCategory not found")
	}

	return trxCategory, nil
}

func (s *trxCategoryService) GetAllTrxCategory() ([]entity.TrxCategory, error) {
	trxCategorys, err := s.trxCategoryRepository.FindAll()

	if err != nil {
		return trxCategorys, err
	}

	if len(trxCategorys) <= 0 {
		return trxCategorys, errors.New("trxCategory not found")
	}

	return trxCategorys, nil
}

func (s *trxCategoryService) DeleteTrxCategory(id int, userName string) (entity.TrxCategory, error) {
	trxCategorys, err := s.GetTrxCategoryById(id)

	if err != nil {
		return entity.TrxCategory{}, err
	}

	if len(trxCategorys) == 0 {
		return entity.TrxCategory{}, errors.New("TrxCategory Not Found")
	}

	trxCategory := trxCategorys[0]

	trxCategory.DeletedDate = time.Now()
	trxCategory.DeletedBy = userName
	result, err := s.trxCategoryRepository.Edit(trxCategory)
	if err != nil {
		return result, err
	}
	return result, nil
}
