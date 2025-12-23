package connection

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Recipient struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
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
