package main

import (
	"github.com/ockibagusp/hello/db"
	t "github.com/ockibagusp/hello/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	c "github.com/ockibagusp/hello/controllers"
)

func main() {
	// PROD or DEV
	dbManager := db.Init("PROD")

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
	controllers := c.Controller{DB: dbManager}

	// Route => controllers
	e.GET("/", controllers.Home).Name = "home"
	e.GET("/about", controllers.About).Name = "about"
	e.GET("/users", controllers.Users).Name = "users"
	e.GET("/users/add", controllers.CreateUser).Name = "user/add get"
	e.POST("/users/add", controllers.CreateUser).Name = "user/add post"
	e.GET("/users/read/:id", controllers.ReadUser).Name = "user/read get"
	e.GET("/users/view/:id", controllers.UpdateUser).Name = "user/view get"
	e.POST("/users/view/:id", controllers.UpdateUser).Name = "user/view post"
	e.GET("/users/delete/:id", controllers.DeleteUser).Name = "user/delete get"

	// Route => controllers API
	g := e.Group("/api/v1")
	g.GET("/users", controllers.UsersAPI).Name = "users get"

	g.POST("/users", controllers.CreateUserAPI).Name = "user post"
	g.GET("/users/:id", controllers.ReadUserAPI).Name = "user read"
	g.PUT("/users/:id", controllers.UpdateUserAPI).Name = "user update"
	g.DELETE("/users/:id", controllers.DeleteUserAPI).Name = "user delete"

	// Start the Echo server
	e.Start(":8000")
}
