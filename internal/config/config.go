package config

import (
	"os"
)

type Config struct {
	HTTPPort string
	MySQLDSN string
}

func Load() Config {
	return Config{
		HTTPPort: os.Getenv("HTTP_PORT"),
		MySQLDSN: os.Getenv("MYSQL_DSN"),
	}
}
