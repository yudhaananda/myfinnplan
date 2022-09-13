package entity

import "time"

type User struct {
	Id          int `gorm:"primarykey;autoIncrement:true"`
	UserName    string
	Password    string
	CreatedBy   string
	CreatedDate time.Time
	UpdatedBy   string
	UpdatedDate time.Time
	DeletedBy   string
	DeletedDate time.Time
}
