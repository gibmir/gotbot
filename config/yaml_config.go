package config

import (
	"fmt"
	"log"
)

const configPathArgName = "-cp"
const defaultConfigPath = "/etc/gotbot/config.yaml"

//
type YamlConfigFactory struct {
	reader *Reader
}

func (factory *YamlConfigFactory) Create() (*Config, error) {
	configPath := factory.reader.DefaultRead(configPathArgName, defaultConfigPath)
	log.Printf("config path [%v] was read", configPath)
	return nil, fmt.Errorf("Unsupported")
}
