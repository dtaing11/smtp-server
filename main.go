package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dtaing11/smtp-server/connection"
	"github.com/dtaing11/smtp-server/db"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
	dbConn, err := db.DbConnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	log.Println("Database connected")

	http.HandleFunc("/sendEmail", connection.ApiKeyMiddleware(connection.EmailSendHandler))

	fmt.Println("Server starting on :8080")
	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
