package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Client struct {
	IP   string
	Port int
	ID   uint32
}

func (c *Client) Parse(path string) {
	if _, err := toml.DecodeFile(path, c); err != nil {
		log.Fatalln(err)
	}
}
