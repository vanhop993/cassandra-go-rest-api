package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routes(router *mux.Router, ch *UserHandlers) {
	router.HandleFunc("/users", ch.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/users", ch.Insert).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", ch.GetById).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", ch.Update).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", ch.Delete).Methods(http.MethodDelete)
}
