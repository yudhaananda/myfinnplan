package service

import (
	"errors"
	"myfinnplan/entity"
	"myfinnplan/input"
	"myfinnplan/repository"
	"time"
)

type TransactionService interface {
	CreateTransaction(input input.TransactionInput, userName string) (entity.Transaction, error)
	EditTransaction(input input.TransactionEditInput, userName string) (entity.Transaction, error)
	GetTransactionById(id int) ([]entity.Transaction, error)
	GetTransactionByBankAccountId(bankAccountId int) ([]entity.Transaction, error)
	GetTransactionByCategoryCode(categoryCode string) ([]entity.Transaction, error)
	GetTransactionByAmount(amount float64) ([]entity.Transaction, error)
	GetTransactionByNotes(notes string) ([]entity.Transaction, error)

	GetAllTransaction() ([]entity.Transaction, error)
	DeleteTransaction(id int, userName string) (entity.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepository}
}

func (s *transactionService) CreateTransaction(input input.TransactionInput, userName string) (entity.Transaction, error) {
	transaction := entity.Transaction{
		BankAccountId: input.BankAccountId,
		Amount:        input.Amount,
		Notes:         input.Notes,
		CreatedBy:     userName,
		CreatedDate:   time.Now(),
	}

	newTransaction, err := s.transactionRepository.Save(transaction)

	if err != nil {
		return transaction, err
	}

	return newTransaction, nil
}

func (s *transactionService) EditTransaction(input input.TransactionEditInput, userName string) (entity.Transaction, error) {
	oldTransactions, err := s.transactionRepository.FindById(input.Id)

	if err != nil {
		return entity.Transaction{}, err
	}

	oldTransaction := oldTransactions[0]

	transaction := entity.Transaction{
		Id:            input.Id,
		BankAccountId: input.BankAccountId,
		Amount:        input.Amount,
		Notes:         input.Notes,
		CreatedBy:     oldTransaction.CreatedBy,
		CreatedDate:   oldTransaction.CreatedDate,
		UpdatedBy:     userName,
	}

	newTransaction, err := s.transactionRepository.Edit(transaction)

	if err != nil {
		return transaction, err
	}

	return newTransaction, nil
}

func (s *transactionService) GetTransactionById(id int) ([]entity.Transaction, error) {

	transaction, err := s.transactionRepository.FindById(id)

	if err != nil {
		return transaction, err
	}

	if len(transaction) == 0 {
		return transaction, errors.New("transaction not found")
	}

	return transaction, nil
}
func (s *transactionService) GetTransactionByBankAccountId(bankAccountId int) ([]entity.Transaction, error) {

	transaction, err := s.transactionRepository.FindByBankAccountId(bankAccountId)

	if err != nil {
		return transaction, err
	}

	if len(transaction) == 0 {
		return transaction, errors.New("transaction not found")
	}

	return transaction, nil
}
func (s *transactionService) GetTransactionByCategoryCode(categoryCode string) ([]entity.Transaction, error) {

	transaction, err := s.transactionRepository.FindByCategoryCode(categoryCode)

	if err != nil {
		return transaction, err
	}

	if len(transaction) == 0 {
		return transaction, errors.New("transaction not found")
	}

	return transaction, nil
}
func (s *transactionService) GetTransactionByAmount(amount float64) ([]entity.Transaction, error) {

	transaction, err := s.transactionRepository.FindByAmount(amount)

	if err != nil {
		return transaction, err
	}

	if len(transaction) == 0 {
		return transaction, errors.New("transaction not found")
	}

	return transaction, nil
}
func (s *transactionService) GetTransactionByNotes(notes string) ([]entity.Transaction, error) {

	transaction, err := s.transactionRepository.FindByNotes(notes)

	if err != nil {
		return transaction, err
	}

	if len(transaction) == 0 {
		return transaction, errors.New("transaction not found")
	}

	return transaction, nil
}

func (s *transactionService) GetAllTransaction() ([]entity.Transaction, error) {
	transactions, err := s.transactionRepository.FindAll()

	if err != nil {
		return transactions, err
	}

	if len(transactions) <= 0 {
		return transactions, errors.New("transaction not found")
	}

	return transactions, nil
}

func (s *transactionService) DeleteTransaction(id int, userName string) (entity.Transaction, error) {
	transactions, err := s.GetTransactionById(id)

	if err != nil {
		return entity.Transaction{}, err
	}

	if len(transactions) == 0 {
		return entity.Transaction{}, errors.New("transaction not found")
	}

	transaction := transactions[0]

	transaction.DeletedDate = time.Now()
	transaction.DeletedBy = userName
	result, err := s.transactionRepository.Edit(transaction)
	if err != nil {
		return result, err
	}
	return result, nil
}
