package test

import (
	"testing"

	dbManager "github.com/ockibagusp/hello/db"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// test database: DEV
var db *gorm.DB = dbManager.Init("DEV")

func TestGetConnectionDB(t *testing.T) {
	assert.NotNil(t, db)
}
