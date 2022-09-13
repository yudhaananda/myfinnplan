package repository

import (
	"myfinnplan/entity"

	"gorm.io/gorm"
)

type BankAccountRepository interface {
	Save(bankAccount entity.BankAccount) (entity.BankAccount, error)
	Edit(bankAccount entity.BankAccount) (entity.BankAccount, error)
	FindById(id int) ([]entity.BankAccount, error)
	FindByAccountCode(accountCode string) ([]entity.BankAccount, error)
	FindByAccountNameOwner(accountNameOwner string) ([]entity.BankAccount, error)
	FindByBankCode(bankCode string) ([]entity.BankAccount, error)
	FindByAmount(amount float64) ([]entity.BankAccount, error)
	FindByNotes(notes string) ([]entity.BankAccount, error)

	FindAll() ([]entity.BankAccount, error)
}

type bankAccountRepository struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) *bankAccountRepository {
	return &bankAccountRepository{db}
}

func (r *bankAccountRepository) Save(bankAccount entity.BankAccount) (entity.BankAccount, error) {
	err := r.db.Create(&bankAccount).Error

	if err != nil {
		return bankAccount, err
	}

	return bankAccount, nil
}

func (r *bankAccountRepository) Edit(bankAccount entity.BankAccount) (entity.BankAccount, error) {
	err := r.db.Save(bankAccount).Error

	if err != nil {
		return bankAccount, err
	}

	return bankAccount, nil
}

func (r *bankAccountRepository) FindById(id int) ([]entity.BankAccount, error) {
	var bankAccount []entity.BankAccount

	err := r.db.Where("id = ? AND deleted_by = ?", id, "").Find(&bankAccount).Error

	if err != nil {
		return bankAccount, err
	}

	return bankAccount, nil
}
func (r *bankAccountRepository) FindByAccountCode(accountCode string) ([]entity.BankAccount, error) {
	var bankAccount []entity.BankAccount

	err := r.db.Where("account_code = ? AND deleted_by = ?", accountCode, "").Find(&bankAccount).Error

	if err != nil {
		return bankAccount, err
	}

	return bankAccount, nil
}
func (r *bankAccountRepository) FindByAccountNameOwner(accountNameOwner string) ([]entity.BankAccount, error) {
	var bankAccount []entity.BankAccount

	err := r.db.Where("account_name_owner = ? AND deleted_by = ?", accountNameOwner, "").Find(&bankAccount).Error

	if err != nil {
		return bankAccount, err
	}

	return bankAccount, nil
}
func (r *bankAccountRepository) FindByBankCode(bankCode string) ([]entity.BankAccount, error) {
	var bankAccount []entity.BankAccount

	err := r.db.Where("bank_code = ? AND deleted_by = ?", bankCode, "").Find(&bankAccount).Error

	if err != nil {
		return bankAccount, err
	}

	return bankAccount, nil
}
func (r *bankAccountRepository) FindByAmount(amount float64) ([]entity.BankAccount, error) {
	var bankAccount []entity.BankAccount

	err := r.db.Where("amount = ? AND deleted_by = ?", amount, "").Find(&bankAccount).Error

	if err != nil {
		return bankAccount, err
	}

	return bankAccount, nil
}
func (r *bankAccountRepository) FindByNotes(notes string) ([]entity.BankAccount, error) {
	var bankAccount []entity.BankAccount

	err := r.db.Where("notes = ? AND deleted_by = ?", notes, "").Find(&bankAccount).Error

	if err != nil {
		return bankAccount, err
	}

	return bankAccount, nil
}

func (r *bankAccountRepository) FindAll() ([]entity.BankAccount, error) {
	var bankAccounts []entity.BankAccount

	err := r.db.Where("deleted_by = ?", "").Find(&bankAccounts).Error

	if err != nil {
		return bankAccounts, err
	}

	return bankAccounts, nil
}
