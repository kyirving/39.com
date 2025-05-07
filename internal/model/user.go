package model

import (
	"time"
)

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(100);null"`
	Phone    string `gorm:"type:varchar(20);null"`
}

func NewUserModel() *User {
	return &User{
		BaseModel: BaseModel{
			Status: 1,
			Ctime:  time.Now(),
			Uptime: time.Now(),
		},
	}
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Add() error {
	return u.GetDB().Table(u.TableName()).Create(u).Error
}
