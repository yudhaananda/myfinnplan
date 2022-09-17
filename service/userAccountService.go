package service

import (
	"errors"
	"myfinnplan/entity"
	"myfinnplan/input"
	"myfinnplan/repository"
	"time"
)

type UserAccountService interface {
	CreateUserAccount(input input.UserAccountInput, userName string) (entity.UserAccount, error)
	EditUserAccount(input input.UserAccountEditInput, userName string) (entity.UserAccount, error)
	GetUserAccountById(id int) ([]entity.UserAccount, error)
	GetUserAccountByAccountCode(accountCode string) ([]entity.UserAccount, error)
	GetUserAccountByAccountName(accountName string) ([]entity.UserAccount, error)

	GetAllUserAccount() ([]entity.UserAccount, error)
	DeleteUserAccount(id int, userName string) (entity.UserAccount, error)
}

type userAccountService struct {
	userAccountRepository repository.UserAccountRepository
}

func NewUserAccountService(userAccountRepository repository.UserAccountRepository) *userAccountService {
	return &userAccountService{userAccountRepository}
}

func (s *userAccountService) CreateUserAccount(input input.UserAccountInput, userName string) (entity.UserAccount, error) {
	userAccount := entity.UserAccount{
		AccountName: input.AccountName,
		CreatedBy:   userName,
		CreatedDate: time.Now(),
	}

	newUserAccount, err := s.userAccountRepository.Save(userAccount)

	if err != nil {
		return userAccount, err
	}

	return newUserAccount, nil
}

func (s *userAccountService) EditUserAccount(input input.UserAccountEditInput, userName string) (entity.UserAccount, error) {
	oldUserAccounts, err := s.userAccountRepository.FindById(input.Id)

	if err != nil {
		return entity.UserAccount{}, err
	}

	oldUserAccount := oldUserAccounts[0]

	userAccount := entity.UserAccount{
		Id:          input.Id,
		AccountName: input.AccountName,
		CreatedBy:   oldUserAccount.CreatedBy,
		CreatedDate: oldUserAccount.CreatedDate,
		UpdatedBy:   userName,
	}

	newUserAccount, err := s.userAccountRepository.Edit(userAccount)

	if err != nil {
		return userAccount, err
	}

	return newUserAccount, nil
}

func (s *userAccountService) GetUserAccountById(id int) ([]entity.UserAccount, error) {

	userAccount, err := s.userAccountRepository.FindById(id)

	if err != nil {
		return userAccount, err
	}

	if len(userAccount) == 0 {
		return userAccount, errors.New("userAccount not found")
	}

	return userAccount, nil
}
func (s *userAccountService) GetUserAccountByAccountCode(accountCode string) ([]entity.UserAccount, error) {

	userAccount, err := s.userAccountRepository.FindByAccountCode(accountCode)

	if err != nil {
		return userAccount, err
	}

	if len(userAccount) == 0 {
		return userAccount, errors.New("userAccount not found")
	}

	return userAccount, nil
}
func (s *userAccountService) GetUserAccountByAccountName(accountName string) ([]entity.UserAccount, error) {

	userAccount, err := s.userAccountRepository.FindByAccountName(accountName)

	if err != nil {
		return userAccount, err
	}

	if len(userAccount) == 0 {
		return userAccount, errors.New("userAccount not found")
	}

	return userAccount, nil
}

func (s *userAccountService) GetAllUserAccount() ([]entity.UserAccount, error) {
	userAccounts, err := s.userAccountRepository.FindAll()

	if err != nil {
		return userAccounts, err
	}

	if len(userAccounts) <= 0 {
		return userAccounts, errors.New("userAccount not found")
	}

	return userAccounts, nil
}

func (s *userAccountService) DeleteUserAccount(id int, userName string) (entity.UserAccount, error) {
	userAccounts, err := s.GetUserAccountById(id)

	if err != nil {
		return entity.UserAccount{}, err
	}

	if len(userAccounts) == 0 {
		return entity.UserAccount{}, errors.New("UserAccount Not Found")
	}

	userAccount := userAccounts[0]

	userAccount.DeletedDate = time.Now()
	userAccount.DeletedBy = userName
	result, err := s.userAccountRepository.Edit(userAccount)
	if err != nil {
		return result, err
	}
	return result, nil
}
