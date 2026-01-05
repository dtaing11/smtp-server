package auth

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthBody struct {
	email    string
	password string
}

func SignUpHandler(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req AuthBody

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Failed Parse Json", http.StatusBadRequest)
		}
		hashPassword, err := passwordHashing(req.password)
		if err != nil {
			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
			return
		}

		hashApiKey, err := generateAPIKey(32)
		if err != nil {
			http.Error(w, "Server Unable to Generate API KEY", http.StatusInternalServerError)
		}

		var userID string
		err = DB.QueryRowContext(r.Context(),
			`insert into app_users (email, password_hash, api_key_hash)
				values ($1, $2, $3)
				returning id`,
			AuthBody.email,
			hashPassword,
			hashApiKey,
		).Scan(&userID)

		if err != nil {
			http.Error(w, "Email is already existed", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"userId": userID,
		})

	}
}

func LoginHandler(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var req AuthBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Failed Parse Json", http.StatusBadRequest)
			return
		}
		if req.email == "" && isValidEmailParse(req.email) && req.password == "" {
			http.Error(w, "Email is empty or invalid or password is empty", http.StatusBadRequest)
		}

		var userID string
		var password_hash string

		err := DB.QueryRowContext(r.Context(),
			`select password_hash, id
			where email = $1`, req.email).Scan(&password_hash, &userID)

		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(req.password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

	}
}
