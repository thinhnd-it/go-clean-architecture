package main

import (
	"fmt"
	"go-clean-architecture/database"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	conn, err := database.ConnectToMySQL()
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %v", err)
	}

	fmt.Println(conn, err)

	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started at http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
