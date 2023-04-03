package controllers

import (
	"encoding/json"
	"net/http"

	models "notification-api/src/models"
	services "notification-api/src/services"

	"github.com/olahol/melody"
)

func UserController(w http.ResponseWriter, r *http.Request, m *melody.Melody) {
	if r.Method == "POST" {
		createUser(w, r, m)
		return
	} else if r.Method == "GET" {
		getUsers(w, r)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func createUser(w http.ResponseWriter, r *http.Request, m *melody.Melody) {
	var user models.User
	error := json.NewDecoder(r.Body).Decode(&user)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	error = services.CreateUser(user)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	error = services.EmitUserCreatedEvent(user, m)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, error := services.GetUsers()
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	response := make(map[string][]*models.User)
	response["data"] = users

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
