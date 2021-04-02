package config

import (
	"flag"
	"fmt"
)

const tokenDefault = "token is not specified"
const configPathDefault = "empty config path"
const commandLineConfigType = "command"
const yamlConfigType = "yaml"

// Represents configuration for bot
type Config struct {
	Token string
}

// Represents configuration factory
type ConfigFactory interface {
	// Creates configuration
	Create() (*Config, error)
}

type CommandLineConfigFactory struct {
	token string
}

// configuration factory that uses flag lib
func (factory CommandLineConfigFactory) Create() (*Config, error) {

	if factory.token == tokenDefault {
		return nil, fmt.Errorf("token wasn't specified")
	}
	return &Config{Token: factory.token} /*success*/, nil /*error*/
}

//
type YamlConfigFactory struct {
	configPath string
}

func (factory YamlConfigFactory) Create() (*Config, error) {
	return nil, fmt.Errorf("Unsupported")
}

func CreateFactory() (ConfigFactory, error) {
	var configType string
	flag.StringVar(&configType, "ct", commandLineConfigType,
		fmt.Sprintf("Bot configuration type {%s, %s}",
			commandLineConfigType, yamlConfigType))
	var token string
	flag.StringVar(&token, "t", tokenDefault, "Bot Token")
	var configPath string
	flag.StringVar(&configPath, "cp", configPathDefault,
		"Path to yaml configuration.")
	flag.Parse()
	if commandLineConfigType == configType {
		return CommandLineConfigFactory{token}, nil
	} else if yamlConfigType == configType {
		if configPathDefault == configPath {
			return nil, fmt.Errorf("you must specify path to the yaml configuration file")
		}
		return YamlConfigFactory{configPath: configPath}, nil
	} else {
		return nil, fmt.Errorf("%s is incorrect. Must be {%s,%s}",
			configType, commandLineConfigType, yamlConfigType)
	}
}
