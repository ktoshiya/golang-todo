package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port string `env:"PORT" envDefault:"dev"`
}

func New() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
