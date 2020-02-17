package controller

import (
	"net/http"
	"strconv"

	"github.com/OckiFals/hello/model"
	"github.com/labstack/echo"
)

// Users ?
func Users(c echo.Context) error {
	return c.Render(http.StatusOK, "user-all.html", map[string]interface{}{
		"name": "Users",
		"nav":  "users", // (?)
	})
}

// CreateUser ?
func CreateUser(c echo.Context) error {
	if "POST" == c.Request().Method {
		u := &model.User{
			ID:       model.Seq,
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Name:     c.FormValue("name"),
		}
		if err := c.Bind(u); err != nil {
			return err
		}
		model.Users[u.ID] = u
		model.Seq++
		return c.JSON(http.StatusOK, u)
	}
	return c.Render(http.StatusOK, "user-add.html", map[string]interface{}{
		"name": "User Add",
		"nav":  "user Add", // (?)
	})
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, model.Users[id])
}

func updateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	model.Users[id].Name = u.Name
	return c.JSON(http.StatusOK, model.Users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(model.Users, id)
	return c.NoContent(http.StatusNoContent)
}
