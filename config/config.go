package config

import (
	"fmt"
)

/*configuration types*/
const commandLineConfigType = "command"
const yamlConfigType = "yaml"
const envConfigType = "env"

/*command line args*/
const configTypeArgName = "-ct"

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
func CreateFactory(reader *CommandLineReader) (ConfigFactory, error) {
	configType := reader.DefaultRead(configTypeArgName,
		commandLineConfigType /*default*/)
	if commandLineConfigType == configType {
		return CommandLineConfigFactory{*reader}, nil
	} else if yamlConfigType == configType {
		return YamlConfigFactory{cReader: *reader}, nil
	} else if envConfigType == configType {
		return EnvConfigFactory{}, nil
	} else {
		return nil, fmt.Errorf("%s is incorrect. Must be {%s,%s,%s}",
			configType, commandLineConfigType, yamlConfigType, envConfigType)
	}
}
