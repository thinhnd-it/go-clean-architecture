package main

import (
	"fmt"
	"go-clean-architecture/api/route"
	"go-clean-architecture/bootstrap"
	"go-clean-architecture/module/user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	app := bootstrap.App()

	db := app.DB

	db.AutoMigrate(&model.User{})

	gin := gin.Default()

	route.Setup(db, gin)

	gin.Run()
}
