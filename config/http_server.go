package config

import "fmt"

type HttpServerConfig struct {
	Host string `envconfig:"HOST" default:"127.0.0.1"`
	Port int    `envconfig:"PORT" default:"8080"`
}

func (c HttpServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
