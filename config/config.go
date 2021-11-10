package config

import (
	"github.com/tkanos/gonfig"
)

// Configuration init
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

// GetConfig init
func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.json", &configuration)
	return configuration
}
