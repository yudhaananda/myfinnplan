package entity

import "time"

type BankAccount struct {
	Id             int `gorm:"primarykey;autoIncrement:true"`
	UserAccountId  int
	UserAccount    UserAccount `gorm:"ForeignKey:UserAccountId"`
	AccountIdOwner int
	BankCode       string
	Bank           Bank `gorm:"foreignkey:Code"`
	Amount         float64
	Notes          string
	IsDebit        bool
	Transactions   []Transaction `gorm:"ForeignKey:BankAccountId"`
	ExpiredDate    time.Time
	CreatedBy      string
	CreatedDate    time.Time
	UpdatedBy      string
	UpdatedDate    time.Time
	DeletedBy      string
	DeletedDate    time.Time
}
