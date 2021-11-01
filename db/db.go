package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ockibagusp/hello/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB
var db *sql.DB
var err error

// Init (?)
func Init() {
	configuration := config.GetConfig()
	connectString := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		configuration.DB_USERNAME,
		configuration.DB_PASSWORD,
		configuration.DB_NAME,
	)
	connection, err = gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db, _ = connection.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(10 * time.Minute)
}

// DbManager (?)
func DbManager() *gorm.DB {
	return connection
}
