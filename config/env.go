package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Port: getEnv("PORT", "3000"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
