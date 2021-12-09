package models

import (
	"errors"

	"gorm.io/gorm"
)

// User: struct
type User struct {
	Model
	Username string `gorm:"unique;not null" form:"username"`
	Email    string `gorm:"unique;not null" form:"email"`
	Password string `gorm:"not null" form:"password"`
	Name     string `gorm:"not null" form:"name"`
	City     uint   `form:"city"`
	Photo    string `form:"photo"`
}

// UserCity: struct
type UserCity struct {
	User
	CityMassage string `gorm:"index:city_massage" form:"city_massage"`
}

/*
 * the tx *DB: exists apparetly
 */

// User: Save
func (user User) Save(db *gorm.DB) (User, error) {
	if err := db.Create(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

// User: FindAll
func (User) FindAll(db *gorm.DB) ([]User, error) {
	users := []User{}

	// Limit: 25 ?
	err := db.Limit(25).Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []User{}, errors.New("User Not Found")
		}
		return []User{}, err
	}

	return users, nil
}

// User: FirstByID
func (user User) FirstByID(db *gorm.DB, id int) (User, error) {
	err := db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("User Not Found")
		}
		return User{}, err
	}

	return user, nil
}

// User: FindByCityID
func (user User) FindByCityID(db *gorm.DB, id int) (User, error) {
	err := db.Select("users.*, cities.id as city_id, cities.city as city_massage").
		Joins("left join cities on users.city = cities.id").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("User Not Found")
		}
		return User{}, err
	}

	return user, nil
}

// User: Update
func (user User) Update(db *gorm.DB, id int) (User, error) {
	err := db.Where("id = ?", id).Updates(&User{
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		City:     user.City,
		Photo:    user.Photo,
	}).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// User: Update By ID and Password
func (user User) UpdateByIDandPassword(db *gorm.DB, id int, password string) (err error) {
	if err = db.Where("id = ?", id).Update("password", password).First(&user).Error; err != nil {
		return err
	}

	return
}

// User: Delete
func (user User) Delete(db *gorm.DB, id int) error {
	// if db.Delete(&user, id).Error != nil {}
	if err := db.Delete(&user, id).Error; err != nil {
		// return errors.New("Record Not Found")
		return err
	}

	return nil
}
