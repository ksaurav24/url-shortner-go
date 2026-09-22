package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Env         string
	DatabaseURL string
}

func MustLoad() Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error occured while loading env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is Required")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is Required")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		panic("DATABASE_URL is Required")
	}

	return Config{
		Port:        port,
		Env:         env,
		DatabaseURL: dbUrl,
	}
}
