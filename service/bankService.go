package service

import (
	"encoding/json"
	"myfinnplan/entity"
	"os"
	"sort"
	"strconv"
)

type BankService interface {
	GetBankData() ([]entity.Bank, error)
	GetBankByCode(code string) (entity.Bank, error)
}

type bankService struct {
}

func NewBankService() *bankService {
	return &bankService{}
}

func (s *bankService) GetBankByCode(code string) (entity.Bank, error) {
	data, err := os.ReadFile("bank.json")
	if err != nil {
		return entity.Bank{}, err
	}

	var banks []entity.Bank

	err = json.Unmarshal(data, &banks)

	if err != nil {
		return entity.Bank{}, err
	}

	bank := entity.Bank{}

	for _, value := range banks {
		if value.Code == code {
			bank = value
		}
	}

	return bank, nil

}

func (s *bankService) GetBankData() ([]entity.Bank, error) {
	data, err := os.ReadFile("bank.json")

	if err != nil {
		return []entity.Bank{}, err
	}

	var bank []entity.Bank

	err = json.Unmarshal(data, &bank)

	sort.Slice(bank, func(i, j int) bool {
		banki, err := strconv.Atoi(bank[i].Code)
		if err != nil {
			return false
		}
		bankj, err := strconv.Atoi(bank[j].Code)
		if err != nil {
			return false
		}
		return banki < bankj
	})

	if err != nil {
		return bank, err
	}

	return bank, nil
}
