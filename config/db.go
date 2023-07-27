package config

import (
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Database     string `envconfig:"DATABASE" default:"telefun"`
	Driver       string `envconfig:"DRIVER" default:"postgres"`
	Host         string `envconfig:"HOST" default:"127.0.0.1"`
	MaxIdleConns int    `envconfig:"MAX_IDLE_CONNS" default:"5"` // minute
	MaxIdleTime  int    `envconfig:"MAX_IDLE_TIME" default:"5"`  // minute
	MaxLifeTime  int    `envconfig:"MAX_LIFE_TIME" default:"4"`  // hour
	MaxOpenConns int    `envconfig:"MAX_OPEN_CONNS" default:"7"`
	Password     string `envconfig:"PASSWORD" default:"postgres"`
	Port         int    `envconfig:"PORT" default:"5432"`
	QueryString  string `envconfig:"QUERY_STRING" default:"sslmode=disable"`
	Username     string `envconfig:"USERNAME" default:"postgres"`
}

func (cfg DBConfig) DSN() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?%s",
		cfg.Driver,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.QueryString,
	)
}
