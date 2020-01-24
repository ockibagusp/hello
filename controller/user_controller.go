package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// CreateUser ?
func CreateUser(c echo.Context) error {
	// if "POST" == c.Request().Method {
	// 	u := &model.User{
	// 		ID: model.Seq,
	// 	}
	// 	if err := c.Bind(u); err != nil {
	// 		return err
	// 	}
	// 	model.Users[u.ID] = u
	// 	model.Seq++
	// }
	return c.Render(http.StatusOK, "user-all.html", map[string]interface{}{
		"name": "Users",
		"nav":  "users", // (?)
		"msg":  "hahaa!",
	})
}
