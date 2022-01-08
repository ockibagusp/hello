package main

import (
	"log"

	"github.com/ockibagusp/hello/controllers"
	"github.com/ockibagusp/hello/router"
)

func main() {
	// controllers init
	controllers := controllers.New()

	// Echo: router
	e := router.New(controllers)

	// Start the Echo server
	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
