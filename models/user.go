package models

import (
	"errors"

	"gorm.io/gorm"
)

// User: struct
type User struct {
	Model
	Username string `gorm:"unique;not null" json:"username" form:"username"`
	Email    string `gorm:"unique;not null" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Name     string `gorm:"not null" json:"name" form:"name"`
	City     uint   `json:"city" form:"city"`
	Photo    string `json:"photo" form:"photo"`
}

// UserCity: struct
type UserCity struct {
	User
	CityMassage string `gorm:"index:city_massage" json:"city_massage" form:"city_massage"`
}

// User: Save
func (user User) Save(db *gorm.DB) (User, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return User{}, err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return User{}, errors.New("User Exists")
	}
	tx.Commit()

	return user, nil
}

// User: FindAll
func (User) FindAll(db *gorm.DB) ([]User, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return []User{}, err
	}

	users := []User{}

	// Limit: 25 ?
	err := tx.Limit(25).Find(&users).Error

	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []User{}, errors.New("User Not Found")
		}
		return []User{}, err
	}
	tx.Commit()

	return users, nil
}

// User: FindByID
func (user User) FindByID(db *gorm.DB, id int) (User, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return User{}, err
	}

	err := tx.First(&user, id).Error

	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("User Not Found")
		}
		return User{}, err
	}
	tx.Commit()

	return user, nil
}

// User: FindByCityID
func (user User) FindByCityID(db *gorm.DB, id int) (User, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return User{}, err
	}

	err := tx.Select("users.*, cities.id as city_id, cities.city as city_massage").
		Joins("left join cities on users.city = cities.id").First(&user, id).Error

	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("User Not Found")
		}
		return User{}, err
	}
	tx.Commit()

	return user, nil
}

// User: Update
func (user User) Update(db *gorm.DB, id int) (User, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return User{}, err
	}

	err := tx.Where("id = ?", id).Updates(&User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		City:     user.City,
		Photo:    user.Photo,
	}).Error

	if err != nil {
		tx.Rollback()
		return User{}, err
	}
	tx.Commit()

	return user, nil
}

// User: Delete
func (user User) Delete(db *gorm.DB, id int) error {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	// if tx.Delete(&user, id).Error != nil {}
	if err := tx.Delete(&user, id).Error; err != nil {
		tx.Rollback()
		// return errors.New("Record Not Found")
		return err
	}
	tx.Commit()

	return nil
}
