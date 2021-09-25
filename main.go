package main

import (
	"github.com/ockibagusp/hello/db"
	t "github.com/ockibagusp/hello/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	c "github.com/ockibagusp/hello/controllers"
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
	e.Static("/assets", "assets")

	// controllers init
	controllers := c.Controller{DB: db.DbManager()}

	// Route => controllers
	e.GET("/", controllers.Home).Name = "home"
	e.GET("/about", controllers.About).Name = "about"
	e.GET("/users", controllers.Users).Name = "users"
	e.GET("/users/add", controllers.CreateUser).Name = "user/add get"
	e.POST("/users/add", controllers.CreateUser).Name = "user/add post"
	e.GET("/users/read/:id", controllers.ReadUser).Name = "user/read get"
	e.GET("/users/view/:id", controllers.UpdateUser).Name = "user/view get"
	e.GET("/users/delete/:id", controllers.DeleteUser).Name = "user/view delete"

	// Start the Echo server
	e.Start(":8000")
}
