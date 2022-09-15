package entity

import "time"

type User struct {
	Id          int `gorm:"primarykey;autoIncrement:true"`
	UserName    string
	Password    string
	Email       string
	Telephone   string
	Photo       string
	IsVerified  bool
	CreatedBy   string
	CreatedDate time.Time
	UpdatedBy   string
	UpdatedDate time.Time
	DeletedBy   string
	DeletedDate time.Time
}
