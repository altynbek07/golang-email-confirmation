package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
	Mail MailConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type MailConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	FromName    string
	FromAddress string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

	return &Config{
		Db: DbConfig{
			Dsn: generateDsn(),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
		Mail: MailConfig{
			Host:        os.Getenv("MAIL_HOST"),
			Port:        port,
			Username:    os.Getenv("MAIL_USERNAME"),
			Password:    os.Getenv("MAIL_PASSWORD"),
			FromName:    os.Getenv("MAIL_FROM_NAME"),
			FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
		},
	}
}

func generateDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"))
}
