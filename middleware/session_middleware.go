package middleware

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ockibagusp/hello/models"
)

// base.html -> {{if eq ((index .session.Values "is_auth_type") | tostring) -1 }}ok{{end}}

// GetUser: get session from User
func GetUser(c echo.Context) (session_gorilla *sessions.Session, err error) {
	if session_gorilla, err = session.Get("session", c); err != nil {
		return
	}

	if _, ok := session_gorilla.Values["username"]; !ok {
		session_gorilla.Values["username"] = ""
	}
	if _, ok := session_gorilla.Values["is_auth_type"]; !ok {
		session_gorilla.Values["is_auth_type"] = -1
	}

	return
}

// SetSession: set session from User
func SetSession(user models.User, c echo.Context) (session_values *SessionValues, err error) {
	session_gorilla, err := session.Get("session", c)
	if err != nil {
		return
	}

	session_gorilla.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days expired
		HttpOnly: true,
		Secure:   true,
	}

	session_gorilla.Values["username"] = user.Username
	// TODO: user.IsAuthType
	session_gorilla.Values["is_auth_type"] = 2 // TODO: admin: 1 and user: 2
	session_gorilla.Save(c.Request(), c.Response())

	if session_values, err = GetSessionValues(session_gorilla.Values); err != nil {
		return nil, err
	}

	return
}

// ClearSession: delete session from User
func ClearSession(c echo.Context) (err error) {
	var session_gorilla *sessions.Session
	if session_gorilla, err = session.Get("session", c); err != nil {
		return
	}

	session_gorilla.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
	}

	session_gorilla.Values["username"] = ""
	session_gorilla.Values["is_auth_type"] = -1
	session_gorilla.Save(c.Request(), c.Response())
	return
}

// RefreshSession: refresh session from User
func RefreshSession(user models.User, c echo.Context) (session_gorilla *sessions.Session, err error) {
	return
}

// -------

type SessionValues struct {
	Username   string
	IsAuthType int
}

// GetSessionValues: get session values from User
func GetSessionValues(session_values_map map[interface{}]interface{}) (
	session_values *SessionValues, err error) {

	// it's session_values_map["is_auth_type"]
	if session_values_map["is_auth_type"] == nil {
		session_values = &SessionValues{
			Username:   "",
			IsAuthType: -1,
		}
		return
	}

	session_values = &SessionValues{
		Username:   session_values_map["username"].(string),
		IsAuthType: session_values_map["is_auth_type"].(int),
	}

	return
}
