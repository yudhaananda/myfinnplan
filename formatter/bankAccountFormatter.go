package formatter

import "myfinnplan/entity"

type BankAccountFormatter struct {
	BankAccount      entity.BankAccount
	TransactionTotal float64
}

func FormatBank(banks []entity.BankAccount) []BankAccountFormatter {
	var bankAccountFormatters []BankAccountFormatter
	if len(banks) > 0 {
		for _, bank := range banks {
			total := 0.0
			if len(bank.Transactions) > 0 {
				for _, value := range bank.Transactions {
					total += value.Amount
				}
			}
			bankAccountFormatters = append(bankAccountFormatters, BankAccountFormatter{BankAccount: bank, TransactionTotal: total})
		}
	}
	return bankAccountFormatters
}
