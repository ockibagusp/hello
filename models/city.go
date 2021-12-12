package models

import (
	"errors"

	"gorm.io/gorm"
)

// City: struct
type City struct {
	ID   uint   `form:"id"`
	City string `form:"city"`
}

// TableName name: string
func (City) TableName() string {
	return "cities"
}

// City: Save
func (city City) Save(db *gorm.DB) (City, error) {
	if err := db.Create(&city).Error; err != nil {
		return City{}, errors.New("City Exists")
	}

	return city, nil
}

// City: FindAll
func (City) FindAll(db *gorm.DB) ([]City, error) {
	cities := []City{}

	if err := db.Find(&cities).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []City{}, errors.New("City Not Found")
		}
		return []City{}, err
	}

	return cities, nil
}

// City: FirstByID
func (city City) FirstByID(db *gorm.DB, id int) (City, error) {
	if err := db.First(&city, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return City{}, errors.New("City Not Found")
		}
		return City{}, err
	}

	return city, nil
}

// City: Delete
func (city City) Delete(db *gorm.DB, id int) error {
	// if db.Delete(&city, id).Error != nil {}
	if err := db.Delete(&city, id).Error; err != nil {
		// return errors.New("Record Not Found")
		return err
	}

	return nil
}
