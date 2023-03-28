package controllers

import (
	"net/http"
	"notification-api/src/configuration"
)

func Routes() {
	m := configuration.NewWS()

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		UserController(w, r, m)
	})

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})
}
