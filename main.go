package main

import (
	t "github.com/OckiFals/hello/template"

	"github.com/OckiFals/hello/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Instantiate a template registry with an array of template set
	e.Renderer = t.Templates()

	// // Why bootstrap.min.css, bootstrap.min.js, jquery.min.js?
	// http.Handle("/", http.FileServer(http.Dir("./assets/css")))
	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// e.Static("/", "assets")

	// Route => controller
	e.GET("/", controller.Home).Name = "home"
	e.GET("/about", controller.About).Name = "about"
	e.GET("/users", controller.Users).Name = "users"
	e.GET("/user/add", controller.CreateUser).Name = "user/add get"
	e.POST("/user/add", controller.CreateUser).Name = "user/add post"

	// Start the Echo server
	e.Start(":8000")
}
