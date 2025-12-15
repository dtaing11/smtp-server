package connection

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"
)

func apiKeyAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expectedKey := os.Getenv("API_KEY")
		if expectedKey == "" {
			log.Println("Server error: API_KEY environment variable not set")
			http.Error(w, "Server error: API_KEY has not yet set up", http.StatusInternalServerError)
			return
		}
		apiKey := r.Header.Get("X-API-Key")
		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedKey)) != 1 {
			http.Error(w, "Unauthroized Access", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
		next(w, r)
	}
}

func protectHandler(w *http.ResponseWriter, r *http.Request) {
	log.Println("Access Granted")

}
