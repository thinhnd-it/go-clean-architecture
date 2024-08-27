package database

import (
	"log"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dbDriver := os.Getenv("DB_CONNECTION")

	switch dbDriver {
	case "mysql":
		DB, err = ConnectToMySQL()
	default:
		log.Fatalf("Unsupported DB_CONNECTION: %s", dbDriver)
	}

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}
