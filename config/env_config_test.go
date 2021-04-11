package config

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestCreateWithError(t *testing.T) {
	factory := EnvConfigFactory{}
	_, err := factory.Create()
	assert.ErrorContains(t, err, TokenEnvName)
}
