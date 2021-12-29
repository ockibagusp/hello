package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/middleware"
	"github.com/ockibagusp/hello/models"
	"github.com/ockibagusp/hello/types"
)

/*
 * Session: Login
 *
 * @target: All
 * @method: GET
 * @route: /login
 */
func (controller *Controller) Login(c echo.Context) error {
	session, err := middleware.GetUser(c)
	if session.Values["is_auth_type"] != -1 && err == nil {
		return c.Redirect(http.StatusFound, "/")
	}

	if c.Request().Method == "POST" {
		passwordForm := &types.LoginForm{
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
		}

		err := passwordForm.Validate()
		if err != nil {
			// TODO: login -> wrong user and password
			return err
		}

		var user models.User
		// err := controller.DB.Select(...).Where(...).Find(...).Error
		if err := controller.DB.Select("username", "password").Where(
			"username = ?", passwordForm.Username,
		).First(&user).Error; err != nil {
			return err
		}

		// check hash password:
		// match = true
		// match = false
		if !middleware.CheckHashPassword(user.Password, passwordForm.Password) {
			return c.Render(http.StatusForbidden, "login.html", echo.Map{
				"is_html_only": true,
			})
		}

		if _, err := middleware.SetSession(user, c); err != nil {
			// err: session not found
			return c.HTML(http.StatusBadRequest, err.Error())
		}

		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "login.html", echo.Map{
		"is_html_only": true,
	})
}

/*
 * Session: Logout
 *
 * @target: Users
 * @method: GET
 * @route: /logout
 */
func (controller *Controller) Logout(c echo.Context) error {
	if err := middleware.ClearSession(c); err != nil {
		// err: session not found
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
