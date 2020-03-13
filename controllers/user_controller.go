package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OckiFals/hello/db"
	"github.com/OckiFals/hello/models"
	"github.com/labstack/echo"
)

// Users ?
func Users(c echo.Context) error {
	db := db.DbManager()

	var users []models.User
	db.Find(&users)

	return c.Render(http.StatusOK, "user-all.html", map[string]interface{}{
		"name":  "Users",
		"nav":   "users", // (?)
		"users": users,
	})
}

// CreateUser ?
func CreateUser(c echo.Context) error {
	db := db.DbManager()
	if "POST" == c.Request().Method {
		var err error
		city64, err := strconv.ParseUint(c.FormValue("city"), 10, 32)
		if err != nil {
			return err
		}
		city := uint(city64)

		user := db.Create(&models.User{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Name:     c.FormValue("name"),
			City:     city,
			Photo:    c.FormValue("photo"),
		})
		if err = c.Bind(user); err != nil {
			return err
		}
		db.FirstOrCreate(&user)

		return c.Redirect(http.StatusMovedPermanently, "/users")
	}

	cities := []models.City{}
	db.Find(&cities)

	return c.Render(http.StatusOK, "user-add.html", map[string]interface{}{
		"name":   "User Add",
		"nav":    "user Add", // (?)
		"cities": cities,
	})
}

// ReadUser (?)
func ReadUser(c echo.Context) error {
	db := db.DbManager()
	// (?)
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.UserCity
	db.Table("users").Select("users.*, cities.id as city_id, cities.city as city_massage").
		Joins("left join cities on users.city = cities.id").
		First(&user, id)

	return c.Render(http.StatusOK, "user-read.html", map[string]interface{}{
		"name": fmt.Sprintf("User: %s", user.Name),
		"nav":  fmt.Sprintf("User: %s", user.Name), // (?)
		"user": user,
	})
}

// UpdateUser (?)
func UpdateUser(c echo.Context) error {
	db := db.DbManager()
	// (?)
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.UserCity
	db.Table("users").Select("users.*, cities.id as city_id, cities.city as city_massage").
		Joins("left join cities on users.city = cities.id").
		First(&user, id)

	cities := []models.City{}
	db.Find(&cities)

	return c.Render(http.StatusOK, "user-view.html", map[string]interface{}{
		"name":   fmt.Sprintf("User: %s", user.Name),
		"nav":    fmt.Sprintf("User: %s", user.Name), // (?)
		"user":   user,
		"cities": cities,
	})
}

func getUser(c echo.Context) error {
	var user models.User
	db := db.DbManager()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, db.Find(&user, id))
}

func updateUser(c echo.Context) error {
	var user models.User
	db := db.DbManager()
	id, _ := strconv.Atoi(c.Param("id"))
	db.Find(&user, id)
	if err := c.Bind(user); err != nil {
		return err
	}
	db.Model(&user).Update("name", "hello")
	return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
	var user models.User
	db := db.DbManager()
	id, _ := strconv.Atoi(c.Param("id"))
	db.Delete(&user, id)
	return c.NoContent(http.StatusNoContent)
}
