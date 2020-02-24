package db

import (
	"fmt"

	"github.com/OckiFals/hello/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

var (
	DB_USERNAME = "root"
	DB_PASSWORD = "AdminPassword"
	DB_NAME     = "hello"
)

// Init (?)
func Init() {
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_NAME)
	db, err = gorm.Open("mysql", connect_string)
	if err != nil {
		panic("failed to connect database")
	}
	// (?)
	// defer db.Close()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.City{})
}

// DbManager (?)
func DbManager() *gorm.DB {
	return db
}
