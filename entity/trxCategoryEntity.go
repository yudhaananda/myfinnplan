package entity

import "time"

type TrxCategory struct {
	Id           int `gorm:"primarykey;autoIncrement:true"`
	CategoryName string
	CreatedBy    string
	CreatedDate  time.Time
	UpdatedBy    string
	UpdatedDate  time.Time
	DeletedBy    string
	DeletedDate  time.Time
}
