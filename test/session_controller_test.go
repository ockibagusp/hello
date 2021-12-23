package test

import (
	"net/http"
	"testing"

	"github.com/ockibagusp/hello/models"
	"github.com/ockibagusp/hello/types"
)

const GET int = 1

// POST int = 2
const POST = 2

func TestLogin(t *testing.T) {
	noAuth := setupTestServer(t)

	// test for db users
	truncateUsers(db)
	models.User{
		Username: "ockibagusp",
		Email:    "ocki.bagus.p@gmail.com",
		Password: "$2a$10$Y3UewQkjw808Ig90OPjuq.zFYIUGgFkWBuYiKzwLK8n3t9S8RYuYa",
		Name:     "Ocki Bagus Pratama",
	}.Save(db)

	testCases := []struct {
		method int
		name   string
		user   types.LoginForm
		status int
	}{
		{
			method: GET,
			name:   "login get",
			// HTTP response status: 200 OK
			status: http.StatusOK,
		},
		{
			method: POST,
			name:   "login success",
			user: types.LoginForm{
				Username: "ockibagusp",
				Password: "user123",
			},
			// HTTP response status: 200 OK
			status: http.StatusOK,
		},
		{
			method: POST,
			name:   "login failure",
			user: types.LoginForm{
				Username: "ockibagusp",
				Password: "<bad password>",
			},
			// HTTP response status: 403 Forbidden
			status: http.StatusForbidden,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.method == GET {
				noAuth.GET("/login").
					Expect().
					Status(tc.status)
				return
			}
			// tc.method == POST
			noAuth.POST("/login").WithForm(tc.user).
				Expect().
				Status(tc.status)
		})
	}
}
