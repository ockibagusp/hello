package controllers

import "gorm.io/gorm"

// Controller is a controller for this application
type Controller struct {
	DB *gorm.DB
}
