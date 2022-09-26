package entity

import "time"

type UserAccount struct {
	Id           int           `gorm:"primarykey;autoIncrement:true"`
	BankAccounts []BankAccount `gorm:"ForeignKey:UserAccountId"`
	AccountName  string
	UserId       int
	CreatedBy    string
	CreatedDate  time.Time
	UpdatedBy    string
	UpdatedDate  time.Time
	DeletedBy    string
	DeletedDate  time.Time
}
