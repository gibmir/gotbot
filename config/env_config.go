package config

import (
	"fmt"
	"os"
)

const tokenEnvName = "TOKEN"

var ErrEnvEmptyToken = fmt.Errorf("[%v] environment variable is null", tokenEnvName)

type EnvConfigFactory struct {
}

func (factory EnvConfigFactory) Create() (*Config, error) {
	
	token := os.Getenv(tokenEnvName)
	if token == "" {
		return nil, ErrEnvEmptyToken
	}
	return &Config{Token: token}, nil
}
