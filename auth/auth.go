package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func SignUpHandler(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			email    string
			password string
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Failed Parse Json", http.StatusBadRequest)
		}
		hashPassword, err := passwordHashing(req.email)
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
			req.email,
			hashPassword,
			hashApiKey,
		).Scan(&userID)

	}
}
