package config

import (
	"fmt"
)

/*configuration types*/
const cmdConfigType = "cmd"
const yamlConfigType = "yaml"
const envConfigType = "env"

/*command line args*/
const configTypeArgName = "-ct"

// Represent config as map
type Reader struct {
	args map[string]string
}

// Provides config param for specified key.
// return empty string if key wasn't present
func (reader Reader) Read(key string) string {
	return reader.args[key]
}

// Provides config param for specified key.
// return specified default value if key wasn't present
func (reader Reader) DefaultRead(key, defaultVal string) string {
	val := reader.args[key]
	if val == "" {
		return defaultVal
	}
	return val
}

// Represents configuration for bot
type Config struct {
	Token string
}

// Represents configuration factory
type ConfigFactory interface {
	// Creates configuration
	Create() (*Config, error)
}

// Creates configuration factory with specified command line reader
func CreateFactory(reader *Reader) (ConfigFactory, error) {
	configType := reader.DefaultRead(configTypeArgName,
		cmdConfigType /*default*/)
	if cmdConfigType == configType {
		return CmdConfigFactory{*reader}, nil
	} else if yamlConfigType == configType {
		return YamlConfigFactory{cReader: *reader}, nil
	} else if envConfigType == configType {
		return EnvConfigFactory{}, nil
	} else {
		return nil, fmt.Errorf("%s is incorrect. Must be {%s,%s,%s}",
			configType, cmdConfigType, yamlConfigType, envConfigType)
	}
}
