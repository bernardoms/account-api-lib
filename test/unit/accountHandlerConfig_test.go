package unit

import (
	accountApiLib "github.com/bernardoms/account-api-lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_create_handler_config(t *testing.T) {
	config := accountApiLib.NewAccountHandleConfig("http://test.com")
	assert.Equal(t, config.URI, "http://test.com")
}
