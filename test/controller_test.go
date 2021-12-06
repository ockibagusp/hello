package test

import (
	"testing"

	"github.com/ockibagusp/hello/controllers"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Controller test
var _controller controllers.Controller = controllers.Controller{
	DB: new(gorm.DB),
}

func TestController(t *testing.T) {
	assert.NotNil(t, _controller.DB)
}
