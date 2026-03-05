package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Server struct {
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

func (c *Server) Parse(path string) {
	if _, err := toml.DecodeFile(path, c); err != nil {
		log.Fatalln(err)
	}
}
