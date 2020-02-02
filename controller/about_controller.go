package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// About "about.html"
func About(c echo.Context) error {
	// Please note the the second parameter "about.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"name": "About",
		"nav":  "about", // (?)
		"msg":  "All about Ocki Bagus Pratama!",
	})
}
