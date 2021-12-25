package test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/ockibagusp/hello/models"
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

	t.Run("users [auth] to GET it success", func(t *testing.T) {
		auth.GET("/users").
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users [no auth] to GET it failure", func(t *testing.T) {
		noAuth.GET("/users").
			Expect().
			// redirect @route: /login
			// HTTP response status: 200 OK
			Status(http.StatusOK)
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

	t.Run("users [no auth] to GET create it success", func(t *testing.T) {
		noAuth.GET("/users/add").
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users [auth] to GET create it success", func(t *testing.T) {
		auth.GET("/users/add").
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("user [auth] to POST create it success", func(t *testing.T) {
		auth.POST("/users/add").
			WithForm(userForm).
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	// Database: " Error 1062: Duplicate entry 'sugriwa@wanara.com' for key 'users.email_UNIQUE' "
	t.Run("users [no auth] to POST create it failure: Duplicate entry", func(t *testing.T) {
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

	// test for db users
	truncateUsers(db)
	models.User{
		Username: "sugriwa",
		Email:    "sugriwa@wanara.com",
		Name:     "Sugriwa",
	}.Save(db)

	t.Run("users [auth] to GET read it success", func(t *testing.T) {
		auth.GET("/users/read/{id}").
			WithPath("id", "1").
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users [auth] to GET read it failure: 1 session and no-id", func(t *testing.T) {
		auth.GET("/users/read/{id}").
			WithPath("id", "-1").
			Expect().
			// HTTP response status: 406 Not Acceptable
			Status(http.StatusNotAcceptable)
	})

	t.Run("users [no auth] to GET read it failure: 2 no-session and id", func(t *testing.T) {
		noAuth.GET("/users/read/{id}").
			WithPath("id", "1").
			Expect().
			// redirect @route: /login
			// HTTP response status: 200 OK
			Status(http.StatusOK)
		// redirection login
	})

	t.Run("users [no auth] to GET read it failure: 3 no-session and no-id", func(t *testing.T) {
		noAuth.GET("/users/read/{id}").
			WithPath("id", "-1").
			Expect().
			// redirect @route: /login
			// HTTP response status: 200 OK
			Status(http.StatusOK)
		// redirection login
	})
}

func TestUpdateUserController(t *testing.T) {
	noAuth := setupTestServer(t)
	auth := setupTestServerAuth(noAuth)

	// test for db users
	truncateUsers(db)
	models.User{
		Username: "subali",
		Email:    "subali@wanara.com",
		Name:     "Subali",
	}.Save(db)

	t.Run("users [auth] to GET update it success", func(t *testing.T) {
		auth.GET("/users/view/{id}").
			WithPath("id", "1").
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users [auth] to POST update it success", func(t *testing.T) {
		auth.POST("/users/view/{id}").
			WithPath("id", "1").
			WithForm(types.UserForm{
				Username: "rahwana",
				Email:    "rahwana@rakshasa.com",
				Name:     "Rahwana",
			}).
			Expect().
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users [auth] to GET update it failure: 1 session and no-id", func(t *testing.T) {
		auth.GET("/users/view/{id}").
			WithPath("id", "-1").
			Expect().
			// HTTP response status: 406 Not Acceptable
			Status(http.StatusNotAcceptable)
	})

	t.Run("users [no auth] to GET update it failure: 2 no-session and id", func(t *testing.T) {
		noAuth.GET("/users/view/{id}").
			WithPath("id", "1").
			Expect().
			// redirect @route: /login
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})

	t.Run("users [no auth] to GET update it failure: 3 no-session and no-id", func(t *testing.T) {
		noAuth.GET("/users/view/{id}").
			WithPath("id", "-1").
			Expect().
			// redirect @route: /login
			// HTTP response status: 200 OK
			Status(http.StatusOK)
	})
}

func TestUpdateUserByPasswordUserController(t *testing.T) {
	noAuth := setupTestServer(t)
	auth := setupTestServerAuth(noAuth)

	// test for db users
	truncateUsers(db)
	users := []models.User{
		{
			Username: "ockibagusp",
			Email:    "ocki.bagus.p@gmail.com",
			Password: "$2a$10$Y3UewQkjw808Ig90OPjuq.zFYIUGgFkWBuYiKzwLK8n3t9S8RYuYa",
			Name:     "Ocki Bagus Pratama",
		},
		{
			Username: "password_success",
			Email:    "password_success@exemple.com",
			Password: "password_success",
			Name:     "Password success",
		},
	}

	for _, user := range users {
		models.User{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Name:     user.Name,
		}.Save(db)
	}

	// t.Run("users [auth] to GET update user by password it success", func(t *testing.T) {
	// 	auth.GET("/users/view/{id}/password").
	// 		WithPath("id", "1").
	// 		Expect().
	// 		// HTTP response status: 200 OK
	// 		Status(http.StatusOK)
	// })

	// t.Run("users [auth] to POST update user by password it success", func(t *testing.T) {
	// 	auth.POST("/users/view/{id}/password").
	// 		WithPath("id", "1").
	// 		WithForm(types.NewPasswordForm{
	// 			OldPassword:        "user123",
	// 			NewPassword:        "password_success",
	// 			ConfirmNewPassword: "password_success",
	// 		}).
	// 		Expect().
	// 		// HTTP response status: 200 OK
	// 		Status(http.StatusOK)
	// })

	// t.Run("users [auth] to POST update user by password it failure: 1"+
	// 	" passwords don't match", func(t *testing.T) {
	// 	auth.POST("/users/view/{id}/password").
	// 		WithPath("id", "1").
	// 		WithForm(types.NewPasswordForm{
	// 			OldPassword:        "user123",
	// 			NewPassword:        "password_success",
	// 			ConfirmNewPassword: "password_failure",
	// 		}).
	// 		Expect().
	// 		// HTTP response status: 403 Forbidden
	// 		Status(http.StatusForbidden)
	// })

	// t.Run("users [auth] to POST update user by password it failure: 2"+
	// 	" username does not matching", func(t *testing.T) {
	// 	auth.POST("/users/view/{id}/password").
	// 		WithPath("id", "2").
	// 		WithForm(types.NewPasswordForm{
	// 			OldPassword:        "user123",
	// 			NewPassword:        "password_failure",
	// 			ConfirmNewPassword: "password_failure",
	// 		}).
	// 		Expect().
	// 		// HTTP response status: 406 Not Acceptable
	// 		Status(http.StatusNotAcceptable)
	// })

	// t.Run("users [no-auth] to GET update user by password it failure: 3"+
	// 	" no session", func(t *testing.T) {
	// 	noAuth.GET("/users/view/{id}/password").
	// 		WithPath("id", "1").
	// 		Expect().
	// 		// redirect @route: /login
	// 		// HTTP response status: 200 OK
	// 		Status(http.StatusOK)
	// })

	// t.Run("users [no-auth] to POST update user by password it failure: 4"+
	// 	" no session", func(t *testing.T) {
	// 	noAuth.POST("/users/view/{id}/password").
	// 		WithPath("id", "1").
	// 		Expect().
	// 		// redirect @route: /login
	// 		// HTTP response status: 200 OK
	// 		Status(http.StatusOK)
	// })

	testCases := []struct {
		name   string
		expect *httpexpect.Expect // auth or no-auth
		method int                // method 1: GET or 2: POST
		path   int                // id:int exemple, id=1
		form   types.NewPasswordForm
		status int
	}{
		{
			name:   "users [auth] to GET update user by password it success",
			expect: auth,
			method: GET,
			path:   1,
			// HTTP response status: 200 OK
			status: http.StatusOK,
		},
		{
			name:   "users [auth] to POST update user by password it success",
			expect: auth,
			method: POST,
			path:   1,
			form: types.NewPasswordForm{
				OldPassword:        "user123",
				NewPassword:        "password_success",
				ConfirmNewPassword: "password_success",
			},
			// HTTP response status: 200 OK
			status: http.StatusOK,
		},
		{
			name: "users [auth] to POST update user by password it failure: 1" +
				" passwords don't match",
			expect: auth,
			method: POST,
			path:   1,
			form: types.NewPasswordForm{
				OldPassword:        "user123",
				NewPassword:        "password_success",
				ConfirmNewPassword: "password_failure",
			},
			// HTTP response status: 403 Forbidden
			status: http.StatusForbidden,
		},
		{
			name: "users [auth] to POST update user by password it failure: 2" +
				" username don't match",
			expect: auth,
			method: POST,
			path:   2,
			form: types.NewPasswordForm{
				OldPassword:        "user123",
				NewPassword:        "password_failure",
				ConfirmNewPassword: "password_failure",
			},
			// HTTP response status: 406 Not Acceptable
			status: http.StatusNotAcceptable,
		},
		{
			name: "users [no-auth] to GET update user by password it failure: 3" +
				" no session",
			expect: noAuth,
			method: GET,
			path:   1,
			// redirect @route: /login
			// HTTP response status: 200 OK
			status: http.StatusOK,
		},
		{
			name: "users [no-auth] to POST update user by password it failure: 4" +
				" no session",
			expect: noAuth,
			method: POST,
			path:   1,
			// redirect @route: /login
			// HTTP response status: 200 OK
			status: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		expect := tc.expect // auth or no-auth

		t.Run(tc.name, func(t *testing.T) {
			if tc.method == GET {
				expect.GET("/users/view/{id}/password").
					WithPath("id", tc.path).
					WithForm(tc.form).
					Expect().
					Status(tc.status)
				return
			} else if tc.method == POST {
				expect.POST("/users/view/{id}/password").
					WithPath("id", tc.path).
					WithForm(tc.form).
					Expect().
					Status(tc.status)
				return
			}
			panic("method: 1=GET and 2=POST")
		})
	}
}
