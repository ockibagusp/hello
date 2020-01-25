package main

import (
	t "github.com/OckiFals/hello/template"

	"github.com/OckiFals/hello/controller"
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

	// Route => controller
	e.GET("/", controller.HomeController).Name = "home"
	e.GET("/about", controller.AboutController).Name = "about"
	e.GET("/users", controller.UserIndexController).Name = "users"

	// Start the Echo server
	e.Start(":8000")
}
