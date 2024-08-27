package main

import (
	"fmt"
	"go-clean-architecture/bootstrap"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	app := bootstrap.App()

	db := app.DB

	fmt.Println(db)

	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
