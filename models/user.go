package models

import (
	"gorm.io/gorm"
)

// User init
type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	City     uint   `json:"city" form:"city"`
	Photo    string `json:"photo" form:"photo"`
}

// UserCity init
type UserCity struct {
	User
	CityMassage string `gorm:"index:city_massage"`
}
