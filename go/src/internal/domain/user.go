package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName string    `gorm:"varchar(255);unique;not null;"`
	Password string    `gorm:"varchar(255);not null;"`
	Email    string    `gorm:"varchar(255);unique;not null;"`
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
