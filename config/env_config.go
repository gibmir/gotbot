package config

import (
	"fmt"
	"os"
)

const TokenEnvName = "TOKEN"

var ErrEnvEmptyToken = fmt.Errorf("[%v] environment variable wasn't specified",
	TokenEnvName)

type EnvConfigFactory struct {
}

func (factory EnvConfigFactory) Create() (*Config, error) {

	token := os.Getenv(TokenEnvName)
	if token == "" {
		return nil, ErrEnvEmptyToken
	}
	return &Config{Token: token}, nil
}
