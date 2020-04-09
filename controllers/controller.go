package controllers

import "github.com/jinzhu/gorm"

// Controller is a controller for this application
type Controller struct {
	DB *gorm.DB
}
