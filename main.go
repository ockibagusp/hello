package main

import (
	"log"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ockibagusp/hello/controllers"
	"github.com/ockibagusp/hello/template"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO: .env cookie store ?
	e.Use(session.Middleware(sessions.NewCookieStore(
		[]byte("something-very-secret"),
	)))

	// Instantiate a template registry with an array of template set
	e.Renderer = template.New()

	// Why bootstrap.min.css, bootstrap.min.js, jquery.min.js?
	e.Static("/assets", "assets")

	// controllers init
	controllers := controllers.New()

	// Route => controllers
	e.GET("/", controllers.Home).Name = "home"
	e.GET("/login", controllers.Login).Name = "login get"
	e.POST("/login", controllers.Login).Name = "login post"
	e.GET("/logout", controllers.Logout).Name = "home"
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
	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
