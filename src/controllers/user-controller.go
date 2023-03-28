package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	config "notification-api/src/configuration"
	models "notification-api/src/models"

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
	parseJsonError := json.NewDecoder(r.Body).Decode(&user)
	if parseJsonError != nil {
		http.Error(w, parseJsonError.Error(), http.StatusBadRequest)
		return
	}

	_, insertDatabaseError := config.ConnectDb().Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if insertDatabaseError != nil {
		http.Error(w, insertDatabaseError.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	emitEventError := emitUserCreatedEvent(user, m)
	if emitEventError != nil {
		http.Error(w, emitEventError.Error(), http.StatusInternalServerError)
		return
	}

	config.ConnectDb().Close()
}

func emitUserCreatedEvent(user models.User, m *melody.Melody) error {
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshaling user: %s\n", err.Error())
		return err
	}

	m.Broadcast(u)

	return nil
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := make([]*models.User, 0)
	rows, databaseError := config.ConnectDb().Query("SELECT name, email FROM users;")
	if databaseError != nil {
		http.Error(w, databaseError.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		user := new(models.User)
		scanError := rows.Scan(&user.Email, &user.Name)
		if scanError != nil {
			http.Error(w, scanError.Error(), http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	response := make(map[string][]*models.User)
	response["data"] = users

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	config.ConnectDb().Close()
}
