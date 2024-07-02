package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/samber/do"
)

type Config struct {
	Port int `env:"PORT"`
}

func NewConfig(i *do.Injector) (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
