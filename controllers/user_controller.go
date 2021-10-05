package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/models"
)

// Users ?
func (controller *Controller) Users(c echo.Context) error {
	var users []models.User
	controller.DB.Limit(25).Find(&users)

	// is parse API: GET /users
	// -> func (controller *...) Users and controller.ParseAPI("/users") ?
	_url := controller.ParseAPI("/users")
	if c.Path() == _url.Path {
		return c.JSON(http.StatusOK, users)
	}

	return c.Render(http.StatusOK, "users/user-all.html", map[string]interface{}{
		"name":  "Users",
		"nav":   "users", // (?)
		"users": users,
	})
}

// CreateUser ?
func (controller *Controller) CreateUser(c echo.Context) error {
	if c.Request().Method == "POST" {
		city64, err := strconv.ParseUint(c.FormValue("city"), 10, 32)
		if err != nil {
			return err
		}
		city := uint(city64)

		user := controller.DB.Create(&models.User{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Name:     c.FormValue("name"),
			City:     city,
			Photo:    c.FormValue("photo"),
		})
		if err := c.Bind(user); err != nil {
			return err
		}
		controller.DB.FirstOrCreate(&user)

		// is parse API: POST /users
		_url := controller.ParseAPI("/users")
		if c.Path() == _url.Path {
			return c.JSON(http.StatusOK, user)
		}

		return c.Redirect(http.StatusMovedPermanently, "/users")
	}

	cities := []models.City{}
	controller.DB.Find(&cities)

	return c.Render(http.StatusOK, "users/user-add.html", map[string]interface{}{
		"name":   "User Add",
		"nav":    "user Add", // (?)
		"cities": cities,
		"is_new": true,
	})
}

// ReadUser ?
func (controller *Controller) ReadUser(c echo.Context) error {
	var user models.UserCity
	id, _ := strconv.Atoi(c.Param("id")) // (?)

	controller.DB.Table("users").Select("users.*, cities.id as city_id, cities.city as city_massage").
		Joins("left join cities on users.city = cities.id").
		First(&user, id)

	// is parse API: GET /users/:id
	_url := controller.ParseAPI("/users/" + strconv.Itoa(id))
	if c.Path() == _url.Path {
		return c.JSON(http.StatusOK, user)
	}

	return c.Render(http.StatusOK, "users/user-read.html", map[string]interface{}{
		"name":    fmt.Sprintf("User: %s", user.Name),
		"nav":     fmt.Sprintf("User: %s", user.Name), // (?)
		"user":    user,
		"is_read": true,
	})
}

// UpdateUser ?
func (controller *Controller) UpdateUser(c echo.Context) error {
	var user models.UserCity
	id, _ := strconv.Atoi(c.Param("id")) // (?)

	controller.DB.Table("users").Select("users.*, cities.id as city_id, cities.city as city_massage").
		Joins("left join cities on users.city = cities.id").
		First(&user, id)

	cities := []models.City{}
	controller.DB.Find(&cities)

	// is parse API: PUT /users/:id
	_url := controller.ParseAPI("/users/" + strconv.Itoa(id))
	if c.Path() == _url.Path {
		return c.JSON(http.StatusOK, user)
	}

	return c.Render(http.StatusOK, "users/user-view.html", map[string]interface{}{
		"name":   fmt.Sprintf("User: %s", user.Name),
		"nav":    fmt.Sprintf("User: %s", user.Name), // (?)
		"user":   user,
		"cities": cities,
	})
}

// DeleteUser ?
func (controller *Controller) DeleteUser(c echo.Context) error {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id")) // (?)
	controller.DB.Delete(&user, id)

	// is parse API: DELETE /users/:id
	_url := controller.ParseAPI("/users/" + strconv.Itoa(id))
	if c.Path() == _url.Path {
		return c.JSON(http.StatusOK, user)
	}

	return c.Redirect(http.StatusMovedPermanently, "/users")
}
