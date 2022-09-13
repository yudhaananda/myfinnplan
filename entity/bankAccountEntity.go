package entity

import "time"

type BankAccount struct {
	Id               int `gorm:"primarykey;autoIncrement:true"`
	AccountCode      string
	AccountNameOwner string
	BankCode         string
	Amount           float64
	Notes            string
	CreatedBy        string
	CreatedDate      time.Time
	UpdatedBy        string
	UpdatedDate      time.Time
	DeletedBy        string
	DeletedDate      time.Time
}
