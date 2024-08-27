package bootstrap

import (
	"go-clean-architecture/database"

	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB
}

func App() Application {
	app := &Application{}

	LoadEnv()

	database.ConnectDatabase()
	app.DB = database.DB
	return *app
}
