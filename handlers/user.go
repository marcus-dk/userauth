package handlers

import (
	"encoding/json"
	"net/http"
	"userauth/controllers"
	"userauth/models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := controllers.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := models.ParseID(mux.Vars(r)["id"])
	user, err := controllers.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdUser, err := controllers.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    id := models.ParseID(mux.Vars(r)["id"])

    var updatedUser models.User
    err := json.NewDecoder(r.Body).Decode(&updatedUser)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Set the ID from the URL to ensure it's updated correctly
    updatedUser.ID = id

    updatedUser, err = controllers.UpdateUser(id, updatedUser)
    if err != nil {
        switch err {
        case controllers.ErrUserNotFound:
            http.Error(w, "User not found", http.StatusNotFound)
        default:
            http.Error(w, "Failed to update user", http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    id := models.ParseID(mux.Vars(r)["id"])

    err := controllers.DeleteUser(id)
    if err != nil {
        switch err {
        case controllers.ErrUserNotFound:
            http.Error(w, "User not found", http.StatusNotFound)
        default:
            http.Error(w, "Failed to delete user", http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
