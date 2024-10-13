package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName string    `gorm:"unique;not null;varchar(255)"`
	Password string    `gorm:"not null;varchar(255)"`
	Email    string    `gorm:"unique;not null;varchar(255)"`
	Birthday time.Time `gorm:"not null"`
}

func NewUser(userName, password, email string, birthday time.Time) *User {
	return &User{
		UserName: userName,
		Password: password,
		Email:    email,
		Birthday: birthday,
	}
}
