package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/middleware"
	"github.com/ockibagusp/hello/models"
	"github.com/ockibagusp/hello/types"
	log "github.com/sirupsen/logrus"
)

/*
 * Session: Login
 *
 * @target: All
 * @method: GET
 * @route: /login
 */
func (controller *Controller) Login(c echo.Context) error {
	log.Info("start GET [@route: /login]")
	session, err := middleware.GetAuth(c)
	if session.Values["is_auth_type"] != -1 && err == nil {
		log.Warn("to [@route: /] session")
		return c.Redirect(http.StatusFound, "/")
	}

	if c.Request().Method == "POST" {
		log.Info("request method POST [@route: /login]")
		passwordForm := &types.LoginForm{
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
		}

		err := passwordForm.Validate()
		if err != nil {
			// TODO: login -> wrong user and password
			log.Warn("for passwordForm.Validate() not nil [@route: /login]")
			return err
		}

		var user models.User
		// err := controller.DB.Select(...).Where(...).Find(...).Error
		if err := controller.DB.Select("username", "password").Where(
			"username = ?", passwordForm.Username,
		).First(&user).Error; err != nil {
			log.Warn("for database `username` or `password` not nil [@route: /login]")
			return err
		}

		// check hash password:
		// match = true
		// match = false
		if !middleware.CheckHashPassword(user.Password, passwordForm.Password) {
			log.Warn("to check wrong hashed password [@route: /login]")
			return c.Render(http.StatusForbidden, "login.html", echo.Map{
				"is_html_only": true,
			})
		}

		if _, err := middleware.SetSession(user, c); err != nil {
			log.Warn("to middleware.SetSession session not found [@route: /login]")
			// err: session not found
			return c.HTML(http.StatusBadRequest, err.Error())
		}

		log.Info("end POST [@route: /]")
		return c.Redirect(http.StatusFound, "/")
	}

	log.Info("end GET [@route: /login]")
	return c.Render(http.StatusOK, "login.html", echo.Map{
		"csrf":         c.Get("csrf"),
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
	log.Info("start GET [@route: /logout]")
	if err := middleware.ClearSession(c); err != nil {
		log.Warn("to middleware.ClearSession session not found [@route: /logout]")
		// err: session not found
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	log.Info("end redirect [@route: /]")
	return c.Redirect(http.StatusSeeOther, "/")
}
