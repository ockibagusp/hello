package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/models"
	"github.com/thedevsaddam/govalidator"
)

// Users ?
func (controller *Controller) Users(c echo.Context) error {
	users, err := models.User{}.FindAll(controller.DB)
	if err != nil {
		return err
	}

	// is parse API: GET /users
	// -> func (controller *...) Users and controller.ParseAPI("/users") ?
	_url := controller.ParseAPI("/users")
	if c.Path() == _url.Path {
		return c.JSON(http.StatusOK, users)
	}

	return c.Render(http.StatusOK, "users/user-all.html", map[string]interface{}{
		"name":  "Users",
		"nav":   "users", // (?)
		"users": users,
	})
}

// CreateUser ?
func (controller *Controller) CreateUser(c echo.Context) error {
	if c.Request().Method == "POST" {
		city64, err := strconv.ParseUint(c.FormValue("city"), 10, 32)
		if err != nil {
			return err
		}
		city := uint(city64)

		user := models.User{
			Username: c.FormValue("username"),
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Name:     c.FormValue("name"),
			City:     city,
			Photo:    c.FormValue("photo"),
		}

		controller.DB.FirstOrCreate(&user)

		if err := c.Bind(&user); err != nil {
			return err
		}

		return c.Redirect(http.StatusMovedPermanently, "/users")
	}

	cities, err := (models.City{}).FindAll(controller.DB)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "users/user-add.html", map[string]interface{}{
		"name":   "User Add",
		"nav":    "user Add", // (?)
		"cities": cities,
		"is_new": true,
	})
}

// ReadUser ?
func (controller *Controller) ReadUser(c echo.Context) error {
	var user models.User
	var err error
	id, _ := strconv.Atoi(c.Param("id")) // (?)

	if user, err = user.FindByID(controller.DB, id); err != nil {
		return err
	}

	// is parse API: GET /users/:id
	_url := controller.ParseAPI("/users/" + strconv.Itoa(id))

	// ":id" ?
	if strings.Replace(c.Path(), ":id", c.Param("id"), 1) == _url.Path {
		return c.JSON(http.StatusOK, user)
	}

	return c.Render(http.StatusOK, "users/user-read.html", map[string]interface{}{
		"name":    fmt.Sprintf("User: %s", user.Name),
		"nav":     fmt.Sprintf("User: %s", user.Name), // (?)
		"user":    user,
		"is_read": true,
	})
}

// UpdateUser ?
func (controller *Controller) UpdateUser(c echo.Context) error {
	var user models.User
	var err error
	id, _ := strconv.Atoi(c.Param("id")) // (?)

	if c.Request().Method == "POST" {
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if _, err := user.Update(controller.DB, id); err != nil {
			return err
		}

		return c.Redirect(http.StatusMovedPermanently, "/users")
	}

	if user, err = user.FindByID(controller.DB, id); err != nil {
		return err
	}

	// Wow
	cities, err := (models.City{}).FindAll(controller.DB)
	if err != nil {
		return err
	}

	// is parse API: PUT /users/:id
	_url := controller.ParseAPI("/users/" + strconv.Itoa(id))

	// ":id" ?
	if strings.Replace(c.Path(), ":id", c.Param("id"), 1) == _url.Path {
		// ?
		if c.Request().Method == "PUT" {
			return c.JSON(http.StatusOK, user)
		}
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	return c.Render(http.StatusOK, "users/user-view.html", map[string]interface{}{
		"name":   fmt.Sprintf("User: %s", user.Name),
		"nav":    fmt.Sprintf("User: %s", user.Name), // (?)
		"user":   user,
		"cities": cities,
	})
}

// DeleteUser ?
func (controller *Controller) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // (?)

	if err := (models.User{}).Delete(controller.DB, id); err != nil {
		return err
	}

	// is parse API: DELETE /users/:id
	_url := controller.ParseAPI("/users/" + strconv.Itoa(id))
	if c.Path() == _url.Path {
		return c.JSON(http.StatusOK, nil)
	}

	// return c.JSON(http.StatusBadRequest, "ok")
	return c.Redirect(http.StatusMovedPermanently, "/users")
}

// -----------

// UsersAPI: GET Users
func (controller *Controller) UsersAPI(c echo.Context) error {
	user := models.User{}
	users, err := user.FindAll(controller.DB)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

// ReadUserAPI: GET User
func (controller *Controller) ReadUserAPI(c echo.Context) error {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id")) // (?)

	user, err := user.FindByID(controller.DB, id)
	if err != nil {
		// err: User Not Found
		fmt.Println("err: ", err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUserAPI: POST User
func (controller *Controller) CreateUserAPI(c echo.Context) error {
	var user models.User
	//var err error

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fmt.Println("1 -> ", user)

	rules := govalidator.MapData{
		"username": []string{"required", "between:3,12"},
		// "email":    []string{"required", "min:5", "max:24", "email"},
		// "password": []string{"required", "min:5", "max:24"},
		// "name":     []string{"required", "min:3"},
	}

	messages := govalidator.MapData{
		"username": []string{"required:You must provide username", "between:The username field must be between 3 to 12 chars"},
		// "email":    []string{"required:You must provide email", "between:The email field must be between 5 to 24 chars"},
		// "password": []string{"required:You must provide password", "between:The password field must be between 5 to 24 chars"},
		// "name":     []string{"required:You must provide name", "between:The name field must be min 3 chars"},
		// "city":     []string{"between:can or not"},
		// "photo":    []string{"between:can or not"},
	}

	opts := govalidator.Options{
		Request:         c.Request(),
		Data:            &user,
		Rules:           rules,
		Messages:        messages,
		RequiredDefault: true,
	}
	gov := govalidator.New(opts)
	fmt.Println("2 -> ", gov)
	validate := gov.Validate()
	fmt.Println("3 -> ", validate)

	if validate != nil {
		fmt.Println("validate No-Nil")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"validationError": validate})
	}

	fmt.Println("validate Nil")

	// user, err = user.Save(controller.DB)

	// if err != nil {
	// 	return err
	// }

	return c.JSON(http.StatusOK, user)
}

// UpdateUserAPI: PUT User
func (controller *Controller) UpdateUserAPI(c echo.Context) error {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id")) // (?)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if _, err := user.Update(controller.DB, id); err != nil {
		return err
	}

	user, err := user.FindByID(controller.DB, id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser API: Delete User
func (controller *Controller) DeleteUserAPI(c echo.Context) error {
	var user models.User
	var err error
	id, _ := strconv.Atoi(c.Param("id")) // (?)

	if err = user.Delete(controller.DB, id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
