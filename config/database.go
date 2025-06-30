package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "app.db"
	}

	database, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	DB = database

}
