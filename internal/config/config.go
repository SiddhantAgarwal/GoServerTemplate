package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port     string `env:"PORT"      envDefault:"8080"`
	DBDSN    string `env:"DB_DSN"    envDefault:""`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:     getEnv("PORT", "8080"),
		DBDSN:    getEnv("DB_DSN", ""),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

func (c *Config) Addr() string {
	return fmt.Sprintf(":%s", c.Port)
}
