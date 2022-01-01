package main

import (
	"log"

	"github.com/labstack/echo/v4/middleware"
	"github.com/ockibagusp/hello/controllers"
	"github.com/ockibagusp/hello/router"
)

func main() {
	// controllers init
	controllers := controllers.New()

	// Echo: router
	e := router.New(controllers)
	// TODO: CSRF no testing
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		// Optional. Default value "header:X-CSRF-Token".
		// Possible values:
		// - "header:<name>"
		// - "form:<name>"
		// - "query:<name>"
		TokenLookup: "form:X-CSRF-Token",
	}))

	// Start the Echo server
	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
