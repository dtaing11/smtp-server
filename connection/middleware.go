package connection

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"
)

func ApiKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
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
