package controllers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/middleware"
	"github.com/ockibagusp/hello/models"
)

// type credentials: of a username and password
type credentials struct {
	username string
	password string
}

// (type credentials) Validate: of a validate username and password
func (lf credentials) Validate() error {
	return validation.ValidateStruct(&lf,
		validation.Field(&lf.username, validation.Required, validation.Length(4, 15)),
		validation.Field(&lf.password, validation.Required, validation.Length(6, 18)),
	)
}

// Session: GET Login
func (controller *Controller) Login(c echo.Context) error {
	session, err := middleware.GetUser(c)
	if len(session.Values) != 0 && err == nil {
		return c.Redirect(http.StatusFound, "/")
	}

	if c.Request().Method == "POST" {
		credentials := &credentials{
			username: c.FormValue("username"),
			password: c.FormValue("password"),
		}

		err := credentials.Validate()
		if err != nil {
			// TODO: login -> wrong user and password
			return err
		}

		var user models.User
		// err := controller.DB.Select(...).Where(...).Find(...).Error
		if err := controller.DB.Select("username", "password").Where(
			"username = ?", credentials.username,
		).Find(&user).Error; err != nil {
			return err
		}

		// check hash password:
		// match = true
		// match = false
		if !middleware.CheckHashPassword(user.Password, credentials.password) {
			return c.Render(http.StatusForbidden, "login.html", echo.Map{
				"is_html_only": true,
			})
		}

		if _, err := middleware.SetSession(user, c); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "400 Bad Request: " + err.Error(),
			})
		}

		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "login.html", echo.Map{
		"is_html_only": true,
	})
}

// Session: GET Logout
func (controller *Controller) Logout(c echo.Context) error {
	if err := middleware.ClearSession(c); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "400 Bad Request: " + err.Error(),
		})
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
