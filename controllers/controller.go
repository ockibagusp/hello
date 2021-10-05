package controllers

import (
	"net/url"
	"path"

	"gorm.io/gorm"
)

// .env (?)
const API string = "/api/v1"

// Controller is a controller for this application
type Controller struct {
	DB  *gorm.DB
	API string
}

// Controller is parse API
func (controller *Controller) ParseAPI(rawurl string) (_url *url.URL) {
	var err error
	_url, err = url.Parse(API)
	if err != nil {
		panic("invalid url")
	}

	_url.Path = path.Join(_url.Path, "/users")

	_url, err = url.Parse(API)
	if err != nil {
		panic("invalid url")
	}

	_url.Path = path.Join(_url.Path, rawurl)
	return
}
