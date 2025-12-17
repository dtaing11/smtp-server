package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dtaing11/smtp-server/connection"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
	http.HandleFunc("/sendEmail", connection.ApiKeyAuth(connection.EmailSendHandler))

	fmt.Println("Server starting on :8080")
	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
