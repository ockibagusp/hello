package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home "home.html"
func (Controller) Home(c echo.Context) error {
	// Please note the the second parameter "home.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "Home",
		"nav":  "home", // (?)
		"msg":  "Ocki Bagus Pratama!",
	})
}
