package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// AboutHandler "about.html"
func AboutHandler(c echo.Context) error {
	// Please note the the second parameter "about.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"name": "About",
		"msg":  "All about Ocki Bagus Pratama!",
	})
}
