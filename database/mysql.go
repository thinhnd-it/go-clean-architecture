package database

import (
	"fmt"
	"log"
	"os"

	env "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMySQL() (*gorm.DB, error) {
	err := env.Load()

	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	fmt.Println("Successfully connected to MySQL using GORM!")
	return db, nil
}
