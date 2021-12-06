package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/middleware"
)

// About "about.html"
func (Controller) About(c echo.Context) error {
	// Please note the the second parameter "about.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	session, _ := middleware.GetUser(c)
	return c.Render(http.StatusOK, "about.html", echo.Map{
		"name":    "About",
		"nav":     "about", // (?)
		"session": session,
		"msg":     "All about Ocki Bagus Pratama!",
	})
}
