package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"userauth/controllers"
	"userauth/models"
	"userauth/database"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdUser, err := controllers.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := controllers.AuthenticateUser(credentials.Email, credentials.Password)
    if err != nil {
        if err == bcrypt.ErrMismatchedHashAndPassword {
            http.Error(w, "Invalid password", http.StatusUnauthorized)
        } else if err == database.ErrUserNotFound { // You'll need to define this error in your database package
            http.Error(w, "User not found", http.StatusUnauthorized)
        } else {
            http.Error(w, "Authentication failed", http.StatusInternalServerError)
        }
        log.Printf("Login failed for email %s: %v", credentials.Email, err)
        return
    }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
