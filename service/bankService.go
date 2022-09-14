package service

import (
	"encoding/json"
	"myfinnplan/entity"
	"os"
)

type BankService interface {
	GetBankData() ([]entity.Bank, error)
}

type bankService struct {

}

func NewBankService() *bankService{
	return &bankService{}
}

func (s *bankService) GetBankData() ([]entity.Bank, error){
	data, err := os.ReadFile("bank.json")

	if err != nil {
		return []entity.Bank{}, err
	}

	var bank []entity.Bank

	err = json.Unmarshal(data, &bank)

	if err != nil {
		return bank, err
	}

	return bank, nil
}