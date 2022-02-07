package config

import (
	"encoding/json"
	"sync"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var (
	once sync.Once
	cfg  *Config
)

type Config struct {
	ServerHost           string `env:"SERVER_HOST"`
	ServerPort           string `env:"SERVER_PORT"`
	ServerReadTimeout    int    `env:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout   int    `env:"SERVER_WRITE_TIMEOUT"`
	ServerMaxHeaderBytes int    `env:"SERVER_MAX_HEADER_BYTES"`
	PostgresHostName     string `env:"POSTGRES_HOST_NAME"`
	PostgresPort         int    `env:"POSTGRES_PORT"`
	PostgresUsername     string `env:"POSTGRES_USERNAME"`
	PostgresPassword     string `env:"POSTGRES_PASSWORD"`
	PostgresDatabaseName string `env:"POSTGRES_DATABASE_NAME"`
}

func (c *Config) String() string {
	b, _ := json.MarshalIndent(c, "", "    ")
	return string(b)
}

func GetConfig() *Config {
	_ = godotenv.Load()

	once.Do(func() {
		config := Config{}
		if err := env.Parse(&config); err != nil {
			panic(err)
		}

		cfg = &config
	})

	return cfg
}
