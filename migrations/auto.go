package main

import (
	"fmt"
	"go/email-verification/internal/user"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(generateDsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&user.User{})
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
