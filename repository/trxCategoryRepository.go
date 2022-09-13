package repository

import (
	"myfinnplan/entity"

	"gorm.io/gorm"
)

type TrxCategoryRepository interface {
	Save(trxCategory entity.TrxCategory) (entity.TrxCategory, error)
	Edit(trxCategory entity.TrxCategory) (entity.TrxCategory, error)
	FindById(id int) ([]entity.TrxCategory, error)
	FindByCategoryCode(categoryCode string) ([]entity.TrxCategory, error)
	FindByCategoryName(categoryName string) ([]entity.TrxCategory, error)

	FindAll() ([]entity.TrxCategory, error)
}

type trxCategoryRepository struct {
	db *gorm.DB
}

func NewTrxCategoryRepository(db *gorm.DB) *trxCategoryRepository {
	return &trxCategoryRepository{db}
}

func (r *trxCategoryRepository) Save(trxCategory entity.TrxCategory) (entity.TrxCategory, error) {
	err := r.db.Create(&trxCategory).Error

	if err != nil {
		return trxCategory, err
	}

	return trxCategory, nil
}

func (r *trxCategoryRepository) Edit(trxCategory entity.TrxCategory) (entity.TrxCategory, error) {
	err := r.db.Save(trxCategory).Error

	if err != nil {
		return trxCategory, err
	}

	return trxCategory, nil
}

func (r *trxCategoryRepository) FindById(id int) ([]entity.TrxCategory, error) {
	var trxCategory []entity.TrxCategory

	err := r.db.Where("id = ? AND deleted_by = ?", id, "").Find(&trxCategory).Error

	if err != nil {
		return trxCategory, err
	}

	return trxCategory, nil
}
func (r *trxCategoryRepository) FindByCategoryCode(categoryCode string) ([]entity.TrxCategory, error) {
	var trxCategory []entity.TrxCategory

	err := r.db.Where("category_code = ? AND deleted_by = ?", categoryCode, "").Find(&trxCategory).Error

	if err != nil {
		return trxCategory, err
	}

	return trxCategory, nil
}
func (r *trxCategoryRepository) FindByCategoryName(categoryName string) ([]entity.TrxCategory, error) {
	var trxCategory []entity.TrxCategory

	err := r.db.Where("category_name = ? AND deleted_by = ?", categoryName, "").Find(&trxCategory).Error

	if err != nil {
		return trxCategory, err
	}

	return trxCategory, nil
}

func (r *trxCategoryRepository) FindAll() ([]entity.TrxCategory, error) {
	var trxCategorys []entity.TrxCategory

	err := r.db.Where("deleted_by = ?", "").Find(&trxCategorys).Error

	if err != nil {
		return trxCategorys, err
	}

	return trxCategorys, nil
}
