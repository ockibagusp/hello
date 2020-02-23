package controllers

import (
	"net/http"
	"strconv"

	"github.com/OckiFals/hello/db"
	"github.com/OckiFals/hello/models"
	"github.com/labstack/echo"
)

// Users ?
func Users(c echo.Context) error {
	db := db.DbManager()
	users := []models.User{}
	db.Find(&users)

	return c.Render(http.StatusOK, "user-all.html", map[string]interface{}{
		"name":  "Users",
		"nav":   "users", // (?)
		"users": users,
	})
}

// CreateUser ?
func CreateUser(c echo.Context) error {
	if "POST" == c.Request().Method {
		db := db.DbManager()
		user := db.Create(&models.User{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Name:     c.FormValue("name"),
		})
		if err := c.Bind(Users); err != nil {
			return err
		}
		db.FirstOrCreate(&user)
		return c.JSON(http.StatusOK, user)
	}
	return c.Render(http.StatusOK, "user-add.html", map[string]interface{}{
		"name": "User Add",
		"nav":  "user Add", // (?)
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
