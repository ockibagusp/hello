package models

import (
	"github.com/jinzhu/gorm"
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
