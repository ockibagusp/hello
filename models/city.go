package models

import (
	"errors"

	"gorm.io/gorm"
)

// City init
type City struct {
	ID   uint   `json:"id" form:"id"`
	City string `json:"city" form:"city"`
}

// TableName name
func (City) TableName() string {
	return "cities"
}

// City: Save
func (city City) Save(db *gorm.DB) (City, error) {
	err := db.Create(&city).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return City{}, err
	}

	return city, nil
}

// City: FindAll
func (City) FindAll(db *gorm.DB) ([]City, error) {
	var err error
	cities := []City{}

	err = db.Find(&cities).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []City{}, errors.New("City Not Found")
	}

	return cities, nil
}

// City: FindByID
func (city City) FindByID(db *gorm.DB, id int) (City, error) {
	err := db.First(&city, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return City{}, errors.New("City Not Found")
	}

	return city, nil
}

// City: Delete
func (city City) Delete(db *gorm.DB, id int) error {
	err := db.Delete(&city, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("City Not Found")
	}

	return nil
}
