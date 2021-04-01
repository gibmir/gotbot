package config

import (
	"flag"
)

const tokenDefault = ""

type Config struct {
	Token string
}

func Create() Config {
	var token = tokenDefault
	flag.StringVar(&token, "t", tokenDefault, "Bot Token")
	flag.Parse()
	return Config{Token: token}
}
