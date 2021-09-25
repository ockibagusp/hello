package models

import (
	"gorm.io/gorm"
)

// User init
type User struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	City     uint
	Photo    string
}

// UserCity init
type UserCity struct {
	User
	CityMassage string `gorm:"index:city_massage"`
}
