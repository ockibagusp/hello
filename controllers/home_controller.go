package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/middleware"
)

// Home "home.html"
func (Controller) Home(c echo.Context) error {
	// Please note the the second parameter "home.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	session, _ := middleware.GetUser(c)
	session_values, _ := middleware.GetSessionValues(session.Values)
	return c.Render(http.StatusOK, "home.html", echo.Map{
		"name":           "Home",
		"nav":            "home", // (?)
		"session_values": session_values,
		"msg":            "Ocki Bagus Pratama!",
	})
}
