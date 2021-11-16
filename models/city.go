package models

import (
	"errors"

	"gorm.io/gorm"
)

// City: struct
type City struct {
	ID   uint   `json:"id" form:"id"`
	City string `json:"city" form:"city"`
}

// TableName name: string
func (City) TableName() string {
	return "cities"
}

// City: Save
func (city City) Save(db *gorm.DB) (City, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return City{}, err
	}

	err := tx.Create(&city).Error

	if err != nil {
		tx.Rollback()
		return City{}, errors.New("City Exists")
	}
	tx.Commit()

	return city, nil
}

// City: FindAll
func (City) FindAll(db *gorm.DB) ([]City, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return []City{}, err
	}

	var err error
	cities := []City{}

	err = tx.Find(&cities).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return []City{}, errors.New("City Not Found")
	}
	tx.Commit()

	return cities, nil
}

// City: FindByID
func (city City) FindByID(db *gorm.DB, id int) (City, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return City{}, err
	}

	err := tx.First(&city, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return City{}, errors.New("City Not Found")
	}
	tx.Commit()

	return city, nil
}

// City: Delete
func (city City) Delete(db *gorm.DB, id int) error {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if tx.Delete(&city, id).Error != nil {
		tx.Rollback()
		return errors.New("Record Not Found")
	}
	tx.Commit()

	return nil
}
