package main

import (
	"github.com/OckiFals/hello/db"
	t "github.com/OckiFals/hello/template"

	"github.com/OckiFals/hello/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db.Init()

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
	e.GET("/", controllers.Home).Name = "home"
	e.GET("/about", controllers.About).Name = "about"
	e.GET("/users", controllers.Users).Name = "users"
	e.GET("/users/add", controllers.CreateUser).Name = "user/add get"
	e.POST("/users/add", controllers.CreateUser).Name = "user/add post"
	e.GET("/users/read/:id", controllers.ReadUser).Name = "user/read get"

	// Start the Echo server
	e.Start(":8000")
}
