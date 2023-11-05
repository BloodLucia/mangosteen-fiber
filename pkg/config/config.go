package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kalougata/bookkeeping/configs"
)

type Config struct {
	DB *configs.Database
}

func NewConfig() *Config {
	return &Config{
		DB: configs.DatabaseConfig(),
	}
}
