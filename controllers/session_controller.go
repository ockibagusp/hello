package controllers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/middleware"
	"github.com/ockibagusp/hello/models"
)

// type credentials: of a username and password
type credentials struct {
	username string
	password string
}

// (type credentials) Validate: of a validate username and password
func (lf credentials) Validate() error {
	return validation.ValidateStruct(&lf,
		validation.Field(&lf.username, validation.Required, validation.Length(4, 15)),
		validation.Field(&lf.password, validation.Required, validation.Length(6, 18)),
	)
}

// Session: GET Login
func (controller *Controller) Login(c echo.Context) error {
	session, err := middleware.GetUser(c)
	if len(session.Values) != 0 && err == nil {
		return c.Redirect(http.StatusFound, "/")
	}

	if c.Request().Method == "POST" {
		credentials := &credentials{
			username: c.FormValue("username"),
			password: c.FormValue("password"),
		}

		err := credentials.Validate()
		if err != nil {
			// TODO: login -> wrong user and password
			return err
		}

		// Password Hash
		hash, err := middleware.PasswordHash(credentials.password)
		if err != nil {
			return err
		}

		// check hash password:
		// match = true
		// match = false
		if !middleware.CheckHashPassword(hash, credentials.password) {
			return err
		}

		var user models.User
		// err := controller.DB.Select(...).Where(...).Find(...).Error
		if err := controller.DB.Select("username, password").Where(
			"username = ? AND password = ?", credentials.username, hash,
		).Find(&user).Error; err != nil {
			return err
		}

		if _, err := middleware.SetSession(user, c); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "400 Bad Request: " + err.Error(),
			})
		}

		return c.Redirect(http.StatusFound, "/")
	}

	return c.HTML(http.StatusOK, loginFormHTML)
}

// Session: GET Logout
func (controller *Controller) Logout(c echo.Context) error {
	if err := middleware.ClearSession(c); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "400 Bad Request: " + err.Error(),
		})
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

/*
why?

TODO:
- template/template: TemplateRenderer.Render(...), data.(map[string]interface{})["foo"] ?
- template/template: TemplateRenderer.Render(...), if errors.As(...) ?
*/
const loginFormHTML string = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Login</title>

    <link rel="canonical" href="/login">

    <!-- Bootstrap core CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" 
    integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    <link rel="stylesheet" href="/assets/css/signin.css">
  </head>

  <body class="text-center">
    <form class="form-signin" action="/login" method="post">
      <img class="mb-4" src="https://getbootstrap.com/docs/4.0/assets/brand/bootstrap-solid.svg" alt="" width="72" height="72">
      <h1 class="h3 mb-3 font-weight-normal">Please sign in</h1>
      <label for="username" class="sr-only">Usernane</label>
      <input type="username" name="username" class="form-control" placeholder="Usernane" value="ockibagusp" required autofocus>
      <label for="password" class="sr-only">Password</label>
      <input type="password" name="password" class="form-control" placeholder="Password" value="user123" required>
      <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
      <p class="mt-5 mb-3 text-muted">&copy; 2021</p>
    </form>
  </body>
</html>
`
