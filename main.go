package main

import (
	t "github.com/OckiFals/hello/template"

	"github.com/OckiFals/hello/handler"
	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	// Instantiate a template registry with an array of template set
	e.Renderer = t.Templates()

	// // Why bootstrap.min.css, bootstrap.min.js, jquery.min.js?
	// http.Handle("/", http.FileServer(http.Dir("./assets/css")))
	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Route => handler
	e.GET("/", handler.HomeHandler).Name = "home"
	e.GET("/about", handler.AboutHandler).Name = "about"

	// Start the Echo server
	e.Start(":8000")
}
