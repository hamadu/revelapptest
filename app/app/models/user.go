package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Password  string `gorm:"not null"`
}
