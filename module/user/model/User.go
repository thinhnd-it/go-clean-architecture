package model

import "gorm.io/gorm"

type User struct {
	Username string `gorm:"type:varchar(100);unique;uniqueIndex;not null"`
	Email    string `gorm:"type:varchar(100);email;unique;uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	gorm.Model
}

type Profile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
