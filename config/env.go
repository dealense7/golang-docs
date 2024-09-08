package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Envs Config
)

type Config struct {
	DBDriver               string
	DBUsername             string
	DBPassword             string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int
}

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Coudn't load env file!!")
	}

	expirationSeconds, _ := strconv.Atoi(os.Getenv("JWT_EXP_SECONDS"))

	Envs = Config{
		DBDriver:               os.Getenv("DB_DRIVER"),
		DBUsername:             os.Getenv("DB_USERNAME"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBName:                 os.Getenv("DB_NAME"),
		JWTSecret:              os.Getenv("JWT_SECRET_KEY"),
		JWTExpirationInSeconds: expirationSeconds,
	}
}
