package connection

import (
	"crypto/subtle"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Recipient struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
}

func ApiKeyAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expectedKey := os.Getenv("API_KEY")
		if expectedKey == "" {
			log.Println("Server error: API_KEY environment variable not set")
			http.Error(w, "Server error: API_KEY is not set", http.StatusInternalServerError)
			return
		}

		apiKey := r.Header.Get("X-API-Key")
		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedKey)) != 1 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func EmailSendHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Access Granted")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var rec Recipient
	if err := json.Unmarshal(body, &rec); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Basic validation
	if rec.EmailAddress == "" {
		http.Error(w, "emailAddress is required", http.StatusBadRequest)
		return
	}

	if err := sendEmail(rec.EmailAddress, rec.FirstName, rec.LastName); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)

		return
	}
	log.Println(rec.LastName)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "Email Sent"})
}
