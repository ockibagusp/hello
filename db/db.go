package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ockibagusp/hello/config"
	"github.com/ockibagusp/hello/models"
)

var db *gorm.DB
var err error

// Init (?)
func Init() {
	configuration := config.GetConfig()
	connectString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err = gorm.Open("mysql", connectString)
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
