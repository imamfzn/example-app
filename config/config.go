package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type Config struct {
	HttpServer HttpServerConfig `envconfig:"HTTP_SERVER"`
	DB         DBConfig         `envconfig:"DB"`
}

func Load() (Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		if err := gotenv.Load(); err != nil {
			return Config{}, err
		}
	}

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, err
}
