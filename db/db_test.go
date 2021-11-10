package db

import (
	"testing"

	"github.com/ockibagusp/hello/config"
	"github.com/stretchr/testify/assert"
)

func TestGetConnectionDB(t *testing.T) {
	assert := assert.New(t)

	db := config.GetConfig()
	assert.NotNil(db)
}
