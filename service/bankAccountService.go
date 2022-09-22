package service

import (
	"errors"
	"myfinnplan/entity"
	"myfinnplan/input"
	"myfinnplan/repository"
	"time"
)

type BankAccountService interface {
	CreateBankAccount(input input.BankAccountInput, userName string) (entity.BankAccount, error)
	EditBankAccount(input input.BankAccountEditInput, userName string) (entity.BankAccount, error)
	GetBankAccountById(id int) ([]entity.BankAccount, error)
	GetBankAccountByAccountCode(accountCode string) ([]entity.BankAccount, error)
	GetBankAccountByAccountIdOwner(id int) ([]entity.BankAccount, error)
	GetBankAccountByBankCode(bankCode string) ([]entity.BankAccount, error)
	GetBankAccountByAmount(amount float64) ([]entity.BankAccount, error)
	GetBankAccountByNotes(notes string) ([]entity.BankAccount, error)

	GetAllBankAccount() ([]entity.BankAccount, error)
	DeleteBankAccount(id int, userName string) (entity.BankAccount, error)
}

type bankAccountService struct {
	bankAccountRepository repository.BankAccountRepository
}

func NewBankAccountService(bankAccountRepository repository.BankAccountRepository) *bankAccountService {
	return &bankAccountService{bankAccountRepository}
}

func (s *bankAccountService) CreateBankAccount(input input.BankAccountInput, userName string) (entity.BankAccount, error) {
	bankAccount := entity.BankAccount{
		AccountIdOwner: input.AccountIdOwner,
		UserAccountId:  input.UserAccountId,
		BankCode:       input.BankCode,
		Amount:         input.Amount,
		Notes:          input.Notes,
		IsDebit:        input.IsDebit,
		ExpiredDate:    input.ExpiredDate,
		CreatedBy:      userName,
		CreatedDate:    time.Now(),
	}

	newBankAccount, err := s.bankAccountRepository.Save(bankAccount)

	if err != nil {
		return bankAccount, err
	}

	return newBankAccount, nil
}

func (s *bankAccountService) EditBankAccount(input input.BankAccountEditInput, userName string) (entity.BankAccount, error) {
	oldBankAccounts, err := s.bankAccountRepository.FindById(input.Id)

	if err != nil {
		return entity.BankAccount{}, err
	}

	oldBankAccount := oldBankAccounts[0]

	bankAccount := entity.BankAccount{
		Id:             input.Id,
		AccountIdOwner: input.AccountIdOwner,
		UserAccountId:  input.UserAccountId,
		BankCode:       input.BankCode,
		Amount:         input.Amount,
		Notes:          input.Notes,
		IsDebit:        input.IsDebit,
		ExpiredDate:    input.ExpiredDate,
		CreatedBy:      oldBankAccount.CreatedBy,
		CreatedDate:    oldBankAccount.CreatedDate,
		UpdatedBy:      userName,
	}

	newBankAccount, err := s.bankAccountRepository.Edit(bankAccount)

	if err != nil {
		return bankAccount, err
	}

	return newBankAccount, nil
}

func (s *bankAccountService) GetBankAccountById(id int) ([]entity.BankAccount, error) {

	bankAccount, err := s.bankAccountRepository.FindById(id)

	if err != nil {
		return bankAccount, err
	}

	if len(bankAccount) == 0 {
		return bankAccount, errors.New("bankAccount not found")
	}

	return bankAccount, nil
}
func (s *bankAccountService) GetBankAccountByAccountCode(accountCode string) ([]entity.BankAccount, error) {

	bankAccount, err := s.bankAccountRepository.FindByAccountCode(accountCode)

	if err != nil {
		return bankAccount, err
	}

	if len(bankAccount) == 0 {
		return bankAccount, errors.New("bankAccount not found")
	}

	return bankAccount, nil
}
func (s *bankAccountService) GetBankAccountByAccountIdOwner(id int) ([]entity.BankAccount, error) {

	bankAccount, err := s.bankAccountRepository.FindByAccountIdOwner(id)

	if err != nil {
		return bankAccount, err
	}

	if len(bankAccount) == 0 {
		return bankAccount, errors.New("bankAccount not found")
	}

	return bankAccount, nil
}
func (s *bankAccountService) GetBankAccountByBankCode(bankCode string) ([]entity.BankAccount, error) {

	bankAccount, err := s.bankAccountRepository.FindByBankCode(bankCode)

	if err != nil {
		return bankAccount, err
	}

	if len(bankAccount) == 0 {
		return bankAccount, errors.New("bankAccount not found")
	}

	return bankAccount, nil
}
func (s *bankAccountService) GetBankAccountByAmount(amount float64) ([]entity.BankAccount, error) {

	bankAccount, err := s.bankAccountRepository.FindByAmount(amount)

	if err != nil {
		return bankAccount, err
	}

	if len(bankAccount) == 0 {
		return bankAccount, errors.New("bankAccount not found")
	}

	return bankAccount, nil
}
func (s *bankAccountService) GetBankAccountByNotes(notes string) ([]entity.BankAccount, error) {

	bankAccount, err := s.bankAccountRepository.FindByNotes(notes)

	if err != nil {
		return bankAccount, err
	}

	if len(bankAccount) == 0 {
		return bankAccount, errors.New("bankAccount not found")
	}

	return bankAccount, nil
}

func (s *bankAccountService) GetAllBankAccount() ([]entity.BankAccount, error) {
	bankAccounts, err := s.bankAccountRepository.FindAll()

	if err != nil {
		return bankAccounts, err
	}

	if len(bankAccounts) <= 0 {
		return bankAccounts, errors.New("bankAccount not found")
	}

	return bankAccounts, nil
}

func (s *bankAccountService) DeleteBankAccount(id int, userName string) (entity.BankAccount, error) {
	bankAccounts, err := s.GetBankAccountById(id)

	if err != nil {
		return entity.BankAccount{}, err
	}

	if len(bankAccounts) == 0 {
		return entity.BankAccount{}, errors.New("BankAccount Not Found")
	}

	bankAccount := bankAccounts[0]

	bankAccount.DeletedDate = time.Now()
	bankAccount.DeletedBy = userName
	result, err := s.bankAccountRepository.Edit(bankAccount)
	if err != nil {
		return result, err
	}
	return result, nil
}
