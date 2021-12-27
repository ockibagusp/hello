package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/middleware"
	"github.com/ockibagusp/hello/models"
	"github.com/ockibagusp/hello/types"
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
	if session.Values["is_auth_type"] == -1 {
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
		"name":    "Users",
		"nav":     "users", // (?)
		"session": session,
		"users":   users,
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
		var city uint
		if c.FormValue("city") != "" {
			city64, err := strconv.ParseUint(c.FormValue("city"), 10, 32)
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "400 Bad Request: " + err.Error(),
				})
			}
			// Kota dan Keb. ?
			city = uint(city64)
		}

		// userForm: type of a user
		_userForm := types.UserForm{
			Username:        c.FormValue("username"),
			Email:           c.FormValue("email"),
			Password:        c.FormValue("password"),
			ConfirmPassword: c.FormValue("confirm_password"),
			Name:            c.FormValue("name"),
			City:            city,
			Photo:           c.FormValue("photo"),
		}

		// _userForm: Validate of a validate user
		err := validation.Errors{
			"username": validation.Validate(
				_userForm.Username, validation.Required, validation.Length(4, 15),
			),
			"email": validation.Validate(_userForm.Email, validation.Required, is.Email),
			"password": validation.Validate(
				_userForm.Password, validation.Required, validation.Length(6, 18),
				validation.By(types.PasswordEquals(_userForm.ConfirmPassword)),
			),
			"name":  validation.Validate(_userForm.Name, validation.Required),
			"city":  validation.Validate(_userForm.City),
			"photo": validation.Validate(_userForm.Photo),
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

		// Password Hash
		hash, err := middleware.PasswordHash(_userForm.Password)
		if err != nil {
			return err
		}

		user := models.User{
			Username: _userForm.Username,
			Email:    _userForm.Email,
			Password: hash,
			Name:     _userForm.Name,
			City:     _userForm.City,
			Photo:    _userForm.Photo,
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

	session, _ := middleware.GetUser(c)
	return c.Render(http.StatusOK, "users/user-add.html", echo.Map{
		"name":    "User Add",
		"nav":     "user Add", // (?)
		"session": session,
		"cities":  cities,
		"is_new":  true,
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
	if session.Values["is_auth_type"] == -1 {
		return c.Redirect(http.StatusFound, "/login")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	// user, err := user.FirstByID(...): be able
	user, err := user.FirstByID(controller.DB, id)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "406 Not Acceptable: " + err.Error(),
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
		"session": session,
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
	if session.Values["is_auth_type"] == -1 {
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

	// user, err := user.FirstByID(...): be able
	user, err := user.FirstByID(controller.DB, id)
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
		"name":    fmt.Sprintf("User: %s", user.Name),
		"nav":     fmt.Sprintf("User: %s", user.Name), // (?)
		"session": session,
		"user":    user,
		"cities":  cities,
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
	if session.Values["is_auth_type"] == -1 {
		return c.Redirect(http.StatusFound, "/login")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	/*
		for example:
		username ockibagusp update by password 'ockibagusp': ok
		username ockibagusp update by password 'sugriwa': no
	*/
	user, err := models.User{}.FirstByIDAndUsername(
		controller.DB, id, session.Values["username"].(string),
	)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "405 Method Not Allowed: " + err.Error(),
		})
	}

	if c.Request().Method == "POST" {
		// newPasswordForm: type of a password user
		_newPasswordForm := types.NewPasswordForm{
			OldPassword:        c.FormValue("old_password"),
			NewPassword:        c.FormValue("new_password"),
			ConfirmNewPassword: c.FormValue("confirm_new_password"),
		}

		if !middleware.CheckHashPassword(user.Password, _newPasswordForm.OldPassword) {
			return c.Render(http.StatusForbidden, "user-view-password.html", echo.Map{
				"session":      session,
				"name":         fmt.Sprintf("User: %s", user.Name),
				"user":         user,
				"is_html_only": true,
			})
		}

		// _newPasswordForm: Validate of a validate user
		err := validation.Errors{
			"password": validation.Validate(
				_newPasswordForm.NewPassword, validation.Required, validation.Length(6, 18),
				validation.By(types.PasswordEquals(_newPasswordForm.ConfirmNewPassword)),
			),
		}.Filter()
		/* if err = validation.Errors{...}.Filter(); err != nil {
			...
		} why?
		*/
		if err != nil {
			// return c.JSON(http.StatusBadRequest, echo.Map{
			// 	"message": "Passwords Don't Match",
			// })
			return c.Render(http.StatusForbidden, "user-view-password.html", echo.Map{
				"session":      session,
				"name":         fmt.Sprintf("User: %s", user.Name),
				"user":         user,
				"is_html_only": true,
			})
		}

		// Password Hash
		hash, err := middleware.PasswordHash(_newPasswordForm.NewPassword)
		if err != nil {
			return err
		}

		user = models.User{
			Password: hash,
		}

		// err := user.UpdateByIDandPassword(...): be able
		if err := user.UpdateByIDandPassword(controller.DB, id, user.Password); err != nil {
			return c.JSON(http.StatusNotAcceptable, echo.Map{
				"message": "405 Method Not Allowed: " + err.Error(),
			})
		}

		return c.Redirect(http.StatusMovedPermanently, "/users")
	}

	/*
		name (string): "users/user-view-password.html" -> no
			{..,"status":500,"error":"html/template: \"users/user-view-password.html\" is undefined",..}
			why?
		name (string): "user-view-password.html" -> yes
	*/
	return c.Render(http.StatusOK, "user-view-password.html", echo.Map{
		"session":      session,
		"name":         fmt.Sprintf("User: %s", user.Name),
		"user":         user,
		"is_html_only": true,
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
	if session.Values["is_auth_type"] == -1 {
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
