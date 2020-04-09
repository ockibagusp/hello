package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OckiFals/hello/models"
	"github.com/labstack/echo"
)

// Users ?
func (controller *Controller) Users(c echo.Context) error {
	var users []models.User
	controller.DB.Find(&users)
	return c.Render(http.StatusOK, "users/user-all.html", map[string]interface{}{
		"name":  "Users",
		"nav":   "users", // (?)
		"users": users,
	})
}

// CreateUser ?
func (controller *Controller) CreateUser(c echo.Context) error {
	if "POST" == c.Request().Method {
		var err error
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
		if err = c.Bind(user); err != nil {
			return err
		}
		controller.DB.FirstOrCreate(&user)

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
	return c.Redirect(http.StatusMovedPermanently, "/users")
}
