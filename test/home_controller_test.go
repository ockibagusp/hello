package test

import (
	"net/http"
	"testing"
)

func TestHomeController(t *testing.T) {
	noAuth := setupTestServer(t)

	// TODO: home test

	t.Run("home success", func(t *testing.T) {
		auth := setupTestServerAuth(noAuth)
		auth.GET("/").
			Expect().
			Status(http.StatusOK)
	})

	t.Run("home failure", func(t *testing.T) {
		noAuth.GET("/").
			Expect().
			Status(http.StatusOK)
	})
}
