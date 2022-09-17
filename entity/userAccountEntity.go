package entity

import "time"

type UserAccount struct {
	Id          int `gorm:"primarykey;autoIncrement:true"`
	AccountName string
	CreatedBy   string
	CreatedDate time.Time
	UpdatedBy   string
	UpdatedDate time.Time
	DeletedBy   string
	DeletedDate time.Time
}
