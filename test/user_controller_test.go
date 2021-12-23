package test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/ockibagusp/hello/types"
	"gorm.io/gorm"
)

// truncate Users
//
// parameter: db *gorm.DB or not available:
// func truncateUsers() {...}, just the same
func truncateUsers(db *gorm.DB) {
	db.Exec("TRUNCATE users")
}

func TestUsersController(t *testing.T) {
	noAuth := setupTestServer(t)
	auth := setupTestServerAuth(noAuth)

	t.Run("users success", func(t *testing.T) {
		auth.GET("/users").
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users failure", func(t *testing.T) {
		noAuth.GET("/users").
			WithRedirectPolicy(httpexpect.DontFollowRedirects).
			Expect().
			// HTTP response status: 302 Found
			Status(http.StatusFound)
	})
}

func TestCreateUserController(t *testing.T) {
	noAuth := setupTestServer(t)
	auth := setupTestServerAuth(noAuth)

	// test for db users
	truncateUsers(db)

	userForm := types.UserForm{
		Username:        "sugriwa",
		Email:           "sugriwa@wanara.com",
		Name:            "Sugriwa",
		Password:        "user123",
		ConfirmPassword: "user123",
	}

	t.Run("users create success: GET", func(t *testing.T) {
		auth.GET("/users/add").
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users create success: POST", func(t *testing.T) {
		auth.POST("/users/add").
			WithForm(userForm).
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users create failure: GET", func(t *testing.T) {
		noAuth.GET("/users/add").
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users create failure: POST", func(t *testing.T) {
		noAuth.POST("/users/add").
			WithForm(userForm).
			Expect().
			// HTTP response status: 400 Bad Request
			Status(http.StatusBadRequest)
	})
}

func TestReadUserController(t *testing.T) {
	noAuth := setupTestServer(t)
	auth := setupTestServerAuth(noAuth)

	t.Run("users [username] success", func(t *testing.T) {
		auth.GET("/users/read/{id}").
			WithPath("id", "1").
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users [username] failure: 1 session and no-id", func(t *testing.T) {
		auth.GET("/users/read/{id}").
			WithPath("id", "-1").
			WithRedirectPolicy(httpexpect.FollowAllRedirects).
			Expect().
			// HTTP response status: 406 Not Acceptable
			Status(http.StatusNotAcceptable)
	})

	t.Run("users [username] failure: 2 no-session and id", func(t *testing.T) {
		noAuth.GET("/users/read/{id}").
			WithPath("id", "1").
			WithRedirectPolicy(httpexpect.FollowAllRedirects).
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
		// redirection login
	})

	t.Run("users [username] failure: 3 no-session and no-id", func(t *testing.T) {
		noAuth.GET("/users/read/{id}").
			WithPath("id", "-1").
			WithRedirectPolicy(httpexpect.FollowAllRedirects).
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
		// redirection login
	})
}
