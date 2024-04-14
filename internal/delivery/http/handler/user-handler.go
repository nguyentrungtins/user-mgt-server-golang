package handler

import (
	"encoding/json"
	"net/http"
	"user-mgt/internal/app/entity"

	"gorm.io/gorm"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the JSON request body into a Credentials struct
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Fetch user from the database
		var user entity.User
		if err := db.Where("username = ?", creds.Username).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// User not found
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}
			// Error fetching user
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Check password
		if user.Password != creds.Password {
			// Incorrect password
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Authentication successful
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	}
}
