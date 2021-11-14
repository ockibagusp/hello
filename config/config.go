package config

import (
	"github.com/tkanos/gonfig"
)

// Configuration: struct
type Configuration struct {
	PROD struct {
		DB_USERNAME string
		DB_PASSWORD string
		DB_PORT     string
		DB_HOST     string
		DB_NAME     string
	}
	DEV struct {
		DB_USERNAME string
		DB_PASSWORD string
		DB_PORT     string
		DB_HOST     string
		DB_NAME     string
	}
}

// GetConfig: Configuration
func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.json", &configuration)
	// test db_test.go
	gonfig.GetConf("../config/config.json", &configuration)
	return configuration
}
