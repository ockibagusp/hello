package controllers

import (
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Controller test
var _controller Controller = Controller{
	DB:  new(gorm.DB),
	API: "/api/v1", // API test
}

func TestController_ParseAPI(t *testing.T) {
	assert := assert.New(t)

	// Lists of normal case scenarios.
	listTest := map[string]string{
		// ("Case", "Parse API: Input|Output")
		"Test1": "/|/api/v1",
		"Test2": "/home|/api/v1/home",
		"Test3": "/users|/api/v1/users",
		"Test4": "/users/1|/api/v1/users/1",
		// incorrect -> "/users/" equal to "/api/v1/users"
		"Test5": "/users/|/api/v1/users",
		"Test6": "/users/1/|/api/v1/users/1",
	}

	var out *url.URL
	for _, value := range listTest {
		inout := strings.Split(value, "|")
		out = _controller.ParseAPI(inout[0])

		assert.Equal(inout[1], out.Path)
	}
}
