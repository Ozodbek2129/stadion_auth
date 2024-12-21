package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVICE string
	AUTH_SERVICE string
	DB_USER      string
	DB_PASSWORD  string
	DB_NAME      string
	DB_HOST      string
	DB_PORT      int
	RD_HOST      string
	RD_PASSWORD  string
	RD_NAME      int
	SIGNING_KEY  string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", ":1234"))
	config.AUTH_SERVICE = cast.ToString(Coalesce("AUTH_SERVICE", ":2345"))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "macbookpro"))
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "google_docs"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "1111"))
	config.DB_PORT = cast.ToInt(Coalesce("DB_PORT", 5432))
	config.RD_HOST = cast.ToString(Coalesce("RD_HOST", "localhost:6379"))
	config.RD_PASSWORD = cast.ToString(Coalesce("RD_PASSWORD", ""))
	config.RD_NAME = cast.ToInt(Coalesce("RD_NAME", 0))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "GOoGLe_DoCs"))
	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
