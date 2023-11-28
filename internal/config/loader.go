package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("err when load env: %s", err.Error())
	}

	return &Config{
		Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			SSL:  os.Getenv("DB_SSL"),
		},
		Server{
			Host: os.Getenv("SRV_HOST"),
			Port: os.Getenv("SRV_PORT"),
		},
	}
}
