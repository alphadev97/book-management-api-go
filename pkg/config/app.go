package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	godotenv.Load()

	dsn := os.Getenv("PORT")
	if dsn == "" {
		log.Fatal("PORT is not found in the environment")
	}

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic(err)
		log.Fatal("error in db connection", err)
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
