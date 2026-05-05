package config

import (
	"os"
	"github.com/joho/godotenv"
)

type config struct {
	DatabaseURL	string
	Port		string
}

func Load() (*config, error) {
	var err error = godotenv.Load()
	
	var cfg *config = &config {
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port: os.Getenv("PORT"),
	}
	
	return cfg, err
}
