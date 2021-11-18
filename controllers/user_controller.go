package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/models"
)

// Users: GET Users
func (controller *Controller) Users(c echo.Context) error {
	// models.User{} or (models.User{}) or var user models.User or user := models.User{}
	users, err := models.User{}.FindAll(controller.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "404 Not Found: " + err.Error(),
		})
	}

	// is parse API: GET /users
	// -> func (controller *...) Users and controller.ParseAPI("/users")
	_url := controller.ParseAPI("/users")
	if c.Path() == _url.Path {
		return c.JSON(http.StatusOK, users)
	}

	return c.Render(http.StatusOK, "users/user-all.html", echo.Map{
		"name":  "Users",
		"nav":   "users", // (?)
		"users": users,
	})
}

// CreateUser: GET or POST User
func (controller *Controller) CreateUser(c echo.Context) error {
	if c.Request().Method == "POST" {
		city64, err := strconv.ParseUint(c.FormValue("city"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "400 Bad Request: " + err.Error(),
			})
		}
		// Kota dan Keb. ?
		city := uint(city64)

		user := models.User{
			Username: c.FormValue("username"),
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Name:     c.FormValue("name"),
			City:     city,
			Photo:    c.FormValue("photo"),
		}

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "400 Bad Request: " + err.Error(),
			})
		}

		// _, err := user.Save(...): be able
		if _, err := user.Save(controller.DB); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Error 1062: Duplicate entry",
			})
		}

		return c.Redirect(http.StatusMovedPermanently, "/users")
	}

	// models.City{} or (models.City{}) or var city models.City or city := models.City{}
	cities, err := models.City{}.FindAll(controller.DB)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	return c.Render(http.StatusOK, "users/user-add.html", echo.Map{
		"name":   "User Add",
		"nav":    "user Add", // (?)
		"cities": cities,
		"is_new": true,
	})
}

// ReadUser: GET User :id
func (controller *Controller) ReadUser(c echo.Context) error {
	var user models.User
	var err error

	id, _ := strconv.Atoi(c.Param("id"))

	// user, _ = user.FindByID(...): be able
	if user, err = user.FindByID(controller.DB, id); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	// models.City{} or (models.City{}) or var city models.City or city := models.City{}
	cities, err := models.City{}.FindAll(controller.DB)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	// is parse API: GET /users/:id
	_url := controller.ParseAPI("/users/" + strconv.Itoa(id))

	// ":id" ?
	if strings.Replace(c.Path(), ":id", c.Param("id"), 1) == _url.Path {
		return c.JSON(http.StatusOK, user)
	}

	return c.Render(http.StatusOK, "users/user-read.html", echo.Map{
		"name":    fmt.Sprintf("User: %s", user.Name),
		"nav":     fmt.Sprintf("User: %s", user.Name), // (?)
		"user":    user,
		"cities":  cities,
		"is_read": true,
	})
}

// UpdateUser: GET or POST User :id
func (controller *Controller) UpdateUser(c echo.Context) error {
	var user models.User
	var err error

	id, _ := strconv.Atoi(c.Param("id"))

	if c.Request().Method == "POST" {
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "400 Bad Request: " + err.Error(),
			})
		}

		// _, err := user.Update(...): be able
		if _, err := user.Update(controller.DB, id); err != nil {
			return c.JSON(http.StatusNotAcceptable, echo.Map{
				"message": "405 Method Not Allowed: " + err.Error(),
			})
		}

		return c.Redirect(http.StatusMovedPermanently, "/users")
	}

	// user, _ = user.FindByID(...): be able
	if user, err = user.FindByID(controller.DB, id); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	// models.City{} or (models.City{}) or var city models.City or city := models.City{}
	cities, err := models.City{}.FindAll(controller.DB)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	// is parse API: PUT /users/:id
	_url := controller.ParseAPI("/users/" + strconv.Itoa(id))

	// ":id" ?
	if strings.Replace(c.Path(), ":id", c.Param("id"), 1) == _url.Path {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "400 Bad Request: " + err.Error(),
		})
	}

	return c.Render(http.StatusOK, "users/user-view.html", echo.Map{
		"name":   fmt.Sprintf("User: %s", user.Name),
		"nav":    fmt.Sprintf("User: %s", user.Name), // (?)
		"user":   user,
		"cities": cities,
	})
}

// DeleteUser: DELETE User :id
func (controller *Controller) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// (models.User{}) or var user models.User or user := models.User{}
	if err := (models.User{}).Delete(controller.DB, id); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	// is parse API: DELETE /users/:id
	_url := controller.ParseAPI("/users/" + strconv.Itoa(id))
	if c.Path() == _url.Path {
		return c.JSON(http.StatusNoContent, echo.Map{
			"message": "204 No Content",
		})
	}

	return c.Redirect(http.StatusMovedPermanently, "/users")
}

// -----------

// UsersAPI: GET Users
func (controller *Controller) UsersAPI(c echo.Context) error {
	// models.User{} or (models.User{}) or var user models.User or user := models.User{}
	user := models.User{}
	users, err := user.FindAll(controller.DB)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "404 Not Found",
		})
	}

	return c.JSON(http.StatusOK, users)
}

// CreateUserAPI: POST User
func (controller *Controller) CreateUserAPI(c echo.Context) error {
	// models.User{} or (models.User{}) or var user models.User or user := models.User{}
	var user models.User
	var err error

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "400 Bad Request: " + err.Error(),
		})
	}

	// user, err = user.Save(...): be able
	user, err = user.Save(controller.DB)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Error 1062: Duplicate entry",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// ReadUserAPI: GET User
func (controller *Controller) ReadUserAPI(c echo.Context) error {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))

	// user, err := user.FindByID(...): be able
	user, err := user.FindByID(controller.DB, id)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUserAPI: PUT User
func (controller *Controller) UpdateUserAPI(c echo.Context) error {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "400 Bad Request: " + err.Error(),
		})
	}

	// _, err := user.Update(...): be able
	if _, err := user.Update(controller.DB, id); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	// user, err := user.FindByID(...): be able
	user, err := user.FindByID(controller.DB, id)

	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser API: Delete User
func (controller *Controller) DeleteUserAPI(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// (models.User{}) or var user models.User or user := models.User{}
	if err := (models.User{}).Delete(controller.DB, id); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, echo.Map{
		"message": "204 No Content",
	})
}
