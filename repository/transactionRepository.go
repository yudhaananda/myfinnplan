package repository

import (
	"myfinnplan/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction entity.Transaction) (entity.Transaction, error)
	Edit(transaction entity.Transaction) (entity.Transaction, error)
	FindById(id int) ([]entity.Transaction, error)
	FindByBankAccountId(bankAccountId int) ([]entity.Transaction, error)
	FindByCategoryCode(categoryCode string) ([]entity.Transaction, error)
	FindByAmount(amount float64) ([]entity.Transaction, error)
	FindByNotes(notes string) ([]entity.Transaction, error)
	FindByUserId(id int) ([]entity.Transaction, error)
	FindAll() ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Save(transaction entity.Transaction) (entity.Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) Edit(transaction entity.Transaction) (entity.Transaction, error) {
	err := r.db.Save(transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) FindById(id int) ([]entity.Transaction, error) {
	var transaction []entity.Transaction

	err := r.db.Where("id = ? AND deleted_by = ?", id, "").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
func (r *transactionRepository) FindByBankAccountId(bankAccountId int) ([]entity.Transaction, error) {
	var transaction []entity.Transaction

	err := r.db.Where("bank_account_id = ? AND deleted_by = ?", bankAccountId, "").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
func (r *transactionRepository) FindByCategoryCode(categoryCode string) ([]entity.Transaction, error) {
	var transaction []entity.Transaction

	err := r.db.Where("category_code = ? AND deleted_by = ?", categoryCode, "").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
func (r *transactionRepository) FindByAmount(amount float64) ([]entity.Transaction, error) {
	var transaction []entity.Transaction

	err := r.db.Where("amount = ? AND deleted_by = ?", amount, "").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) FindByUserId(id int) ([]entity.Transaction, error) {
	var transaction []entity.Transaction

	err := r.db.Table("transactions").Joins("bank_accounts on transactions.bank_account_id = bank_accounts.id").Where("bank_accounts.user_account_id = ? AND transactions.deleted_by = ?", id, "").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
func (r *transactionRepository) FindByNotes(notes string) ([]entity.Transaction, error) {
	var transaction []entity.Transaction

	err := r.db.Where("notes = ? AND deleted_by = ?", notes, "").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) FindAll() ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := r.db.Where("deleted_by = ?", "").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
