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
	if err := db.Create(&user).Error; err != nil {
		return User{}, errors.New("User Exists")
	}

	return user, nil
}

// User: FindAll
func (User) FindAll(db *gorm.DB) ([]User, error) {
	users := []User{}

	// Limit: 25 ?
	err := db.Limit(25).Find(&users).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []User{}, errors.New("User Not Found")
	}

	return users, nil
}

// User: FindByID
func (user User) FindByID(db *gorm.DB, id int) (User, error) {
	err := db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, errors.New("User Not Found")
	}

	return user, nil
}

// User: FindByCityID
func (user User) FindByCityID(db *gorm.DB, id int) (User, error) {
	err := db.Select("users.*, cities.id as city_id, cities.city as city_massage").
		Joins("left join cities on users.city = cities.id").First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, errors.New("User Not Found")
	}

	return user, nil
}

// User: Update
func (user User) Update(db *gorm.DB, id int) (User, error) {
	err := db.Where("id = ?", id).Updates(&User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		City:     user.City,
		Photo:    user.Photo,
	}).Error

	if err != nil {
		return User{}, errors.New("User Not Found")
	}

	return user, nil
}

// User: Delete
func (user User) Delete(db *gorm.DB, id int) error {
	if db.Delete(&user, id).Error != nil {
		return errors.New("Record Not Found")
	}

	return nil
}
