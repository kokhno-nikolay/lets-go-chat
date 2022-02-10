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
	ServerPort           string `env:"SERVER_PORT" envDefault:"3001"`
	ServerReadTimeout    int    `env:"SERVER_READ_TIMEOUT" `
	ServerWriteTimeout   int    `env:"SERVER_WRITE_TIMEOUT"`
	ServerMaxHeaderBytes int    `env:"SERVER_MAX_HEADER_BYTES"`
	PostgresHostName     string `env:"POSTGRES_HOST_NAME" envDefault:"postgres"`
	PostgresPort         int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUsername     string `env:"POSTGRES_USERNAME" envDefault:"postgres"`
	PostgresPassword     string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	PostgresDatabaseName string `env:"POSTGRES_DATABASE_NAME" envDefault:"lets_go_chat"`
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
