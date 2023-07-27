package config

type HttpServerConfig struct {
	Addr string `envconfig:"ADDR" default:"127.0.0.1:8080"`
}
