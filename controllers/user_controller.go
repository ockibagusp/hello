package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/middleware"
	"github.com/ockibagusp/hello/models"
)

/*
 * Users All
 *
 * @target: Users
 * @method: GET
 * @route: /users
 */
func (controller *Controller) Users(c echo.Context) error {
	session, _ := middleware.GetUser(c)
	if len(session.Values) == 0 {
		// TODO: why? return c.Redirect(...): html "status":500,"error":"invalid redirect status code"
		// return c.Redirect(http.StatusUnauthorized, "/login")
		return c.Redirect(http.StatusFound, "/login")
	}

	// models.User{} or (models.User{}) or var user models.User or user := models.User{}
	users, err := models.User{}.FindAll(controller.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "404 Not Found: " + err.Error(),
		})
	}

	return c.Render(http.StatusOK, "users/user-all.html", echo.Map{
		"name":  "Users",
		"nav":   "users", // (?)
		"users": users,
	})
}

/*
 * User Add
 *
 * @target: Users
 * @method: GET or POST
 * @route: /users/add
 */
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

		// START moves dictionary ?

		// type userForm: of a user
		type userForm struct {
			username         string
			email            string
			password         string
			confirm_password string
			name             string
			city             uint
			photo            string
		}

		// function passwordEquals: of password equals confirm_password
		passwordEquals := func(confirm_password string) validation.RuleFunc {
			return func(value interface{}) error {
				password, _ := value.(string)
				if password != confirm_password {
					return errors.New("unexpected password")
				}
				return nil
			}
		}

		// userForm: type of a user
		_userForm := userForm{
			username:         c.FormValue("username"),
			email:            c.FormValue("email"),
			password:         c.FormValue("password"),
			confirm_password: c.FormValue("confirm_password"),
			name:             c.FormValue("name"),
			city:             city,
			photo:            c.FormValue("photo"),
		}

		// _userForm: Validate: of a validate user
		err = validation.Errors{
			"username": validation.Validate(
				_userForm.username, validation.Required, validation.Length(4, 15),
			),
			"email": validation.Validate(_userForm.email, validation.Required, is.Email),
			"password": validation.Validate(
				_userForm.password, validation.Required, validation.Length(6, 18),
				validation.By(passwordEquals(_userForm.confirm_password)),
			),
			"name":  validation.Validate(_userForm.name, validation.Required),
			"city":  validation.Validate(_userForm.city),
			"photo": validation.Validate(_userForm.photo),
		}.Filter()
		/* if err = validation.Errors{...}.Filter(); err != nil {
			...
		} why?
		*/
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "400 Bad Request: " + err.Error(),
			})
		}

		// END

		// Password Hash
		hash, err := middleware.PasswordHash(_userForm.password)
		if err != nil {
			return err
		}

		// check hash password:
		// match = true
		// match = false
		if !middleware.CheckHashPassword(hash, _userForm.password) {
			return err
		}

		user := models.User{
			Username: _userForm.username,
			Email:    _userForm.email,
			Password: hash,
			Name:     _userForm.name,
			City:     _userForm.city,
			Photo:    _userForm.photo,
		}

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "400 Bad Request: " + err.Error(),
			})
		}

		// _, err := user.Save(...): be able
		if _, err := user.Save(controller.DB); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
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

/*
 * Read User ID
 *
 * @target: Users
 * @method: GET
 * @route: /users/read/:id
 */
func (controller *Controller) ReadUser(c echo.Context) error {
	session, _ := middleware.GetUser(c)
	if len(session.Values) == 0 {
		// return c.Redirect(http.StatusUnauthorized, "/login")
		return c.Redirect(http.StatusFound, "/login")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	// user, err := user.FindByID(...): be able
	user, err := user.FindByID(controller.DB, id)
	if err != nil {
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

	return c.Render(http.StatusOK, "users/user-read.html", echo.Map{
		"name":    fmt.Sprintf("User: %s", user.Name),
		"nav":     fmt.Sprintf("User: %s", user.Name), // (?)
		"user":    user,
		"cities":  cities,
		"is_read": true,
	})
}

/*
 * Update User ID
 *
 * @target: Users
 * @method: GET or POST
 * @route: /users/view/:id
 */
func (controller *Controller) UpdateUser(c echo.Context) error {
	session, _ := middleware.GetUser(c)
	if len(session.Values) == 0 {
		// return c.Redirect(http.StatusUnauthorized, "/login")
		return c.Redirect(http.StatusFound, "/login")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	if c.Request().Method == "POST" {
		// TODO: html flash message
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "400 Bad Request: " + err.Error(),
			})
		}

		// _, err = user.Update(...): be able
		if _, err := user.Update(controller.DB, id); err != nil {
			return c.JSON(http.StatusNotAcceptable, echo.Map{
				"message": "405 Method Not Allowed: " + err.Error(),
			})
		}

		return c.Redirect(http.StatusMovedPermanently, "/users")
	}

	// user, err := user.FindByID(...): be able
	user, err := user.FindByID(controller.DB, id)
	if err != nil {
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

	return c.Render(http.StatusOK, "users/user-view.html", echo.Map{
		"name":   fmt.Sprintf("User: %s", user.Name),
		"nav":    fmt.Sprintf("User: %s", user.Name), // (?)
		"user":   user,
		"cities": cities,
	})
}

/*
 * Update User ID by Password
 *
 * @target: Users
 * @method: GET or POST
 * @route: /users/view/:id/password
 */
func (controller *Controller) UpdateUserByPassword(c echo.Context) error {
	session, _ := middleware.GetUser(c)
	if len(session.Values) == 0 {
		// return c.Redirect(http.StatusUnauthorized, "/login")
		return c.Redirect(http.StatusFound, "/login")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	if c.Request().Method == "POST" {
		if c.FormValue("password") != c.FormValue("confirm_password") {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Passwords Don't Match",
			})
		}

		user := models.User{
			Password: c.FormValue("password"),
		}

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "400 Bad Request: " + err.Error(),
			})
		}

		// err := user.UpdateByIDandPassword(...): be able
		if err := user.UpdateByIDandPassword(controller.DB, id, user.Password); err != nil {
			return c.JSON(http.StatusNotAcceptable, echo.Map{
				"message": "405 Method Not Allowed: " + err.Error(),
			})
		}

		return c.Redirect(http.StatusMovedPermanently, "/users")
	}

	var user models.User
	var err error
	// _, err = user.FindByID(...): be able
	if user, err = user.FindByID(controller.DB, id); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	return c.Render(http.StatusOK, "users/user-view-password.html", echo.Map{
		"name": fmt.Sprintf("User: %s", user.Name),
		"nav":  fmt.Sprintf("User: %s", user.Name), // (?)
		"user": user,
	})
}

/*
 * Delete User ID
 *
 * @target: Users
 * @method: GET
 * @route: /users/delete/:id
 */
func (controller *Controller) DeleteUser(c echo.Context) error {
	session, _ := middleware.GetUser(c)
	if len(session.Values) == 0 {
		// return c.Redirect(http.StatusUnauthorized, "/login")
		return c.Redirect(http.StatusFound, "/login")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	// (models.User{}) or var user models.User or user := models.User{}
	if err := (models.User{}).Delete(controller.DB, id); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	return c.Redirect(http.StatusMovedPermanently, "/users")
}
