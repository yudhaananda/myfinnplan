package repository

import (
	"myfinnplan/entity"

	"gorm.io/gorm"
)

type UserAccountRepository interface {
	Save(userAccount entity.UserAccount) (entity.UserAccount, error)
	Edit(userAccount entity.UserAccount) (entity.UserAccount, error)
	FindById(id int) ([]entity.UserAccount, error)
	FindByAccountCode(accountCode string) ([]entity.UserAccount, error)
	FindByAccountName(accountName string) ([]entity.UserAccount, error)

	FindAll() ([]entity.UserAccount, error)
}

type userAccountRepository struct {
	db *gorm.DB
}

func NewUserAccountRepository(db *gorm.DB) *userAccountRepository {
	return &userAccountRepository{db}
}

func (r *userAccountRepository) Save(userAccount entity.UserAccount) (entity.UserAccount, error) {
	err := r.db.Create(&userAccount).Error

	if err != nil {
		return userAccount, err
	}

	return userAccount, nil
}

func (r *userAccountRepository) Edit(userAccount entity.UserAccount) (entity.UserAccount, error) {
	err := r.db.Save(userAccount).Error

	if err != nil {
		return userAccount, err
	}

	return userAccount, nil
}

func (r *userAccountRepository) FindById(id int) ([]entity.UserAccount, error) {
	var userAccount []entity.UserAccount

	err := r.db.Where("id = ? AND deleted_by = ?", id, "").Find(&userAccount).Error

	if err != nil {
		return userAccount, err
	}

	return userAccount, nil
}
func (r *userAccountRepository) FindByAccountCode(accountCode string) ([]entity.UserAccount, error) {
	var userAccount []entity.UserAccount

	err := r.db.Where("account_code = ? AND deleted_by = ?", accountCode, "").Find(&userAccount).Error

	if err != nil {
		return userAccount, err
	}

	return userAccount, nil
}
func (r *userAccountRepository) FindByAccountName(accountName string) ([]entity.UserAccount, error) {
	var userAccount []entity.UserAccount

	err := r.db.Where("account_name = ? AND deleted_by = ?", accountName, "").Find(&userAccount).Error

	if err != nil {
		return userAccount, err
	}

	return userAccount, nil
}

func (r *userAccountRepository) FindAll() ([]entity.UserAccount, error) {
	var userAccounts []entity.UserAccount

	err := r.db.Where("deleted_by = ?", "").Find(&userAccounts).Error

	if err != nil {
		return userAccounts, err
	}

	return userAccounts, nil
}
