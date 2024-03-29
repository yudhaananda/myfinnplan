package entity

import "time"

type Transaction struct {
	Id            int `gorm:"primarykey;autoIncrement:true"`
	BankAccountId int
	BankAccount   BankAccount `gorm:"ForeignKey:BankAccountId"`
	Amount        float64
	Notes         string
	CreatedBy     string
	CreatedDate   time.Time
	UpdatedBy     string
	UpdatedDate   time.Time
	DeletedBy     string
	DeletedDate   time.Time
}
