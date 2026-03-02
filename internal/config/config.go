package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	IP       string
	Repeater Repeater
	HTTP     HTTP
}

type Repeater struct {
	Port     int
	Callsign string
}

type HTTP struct {
	Port int
}

func Parse(path string) *Config {
	config := new(Config)

	if _, err := toml.DecodeFile(path, config); err != nil {
		log.Fatalln(err)
	}

	return config
}
