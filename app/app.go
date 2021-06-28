package app

import (
	"cassandra_rest_api_users/domain"
	"cassandra_rest_api_users/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	DB, _ := DBConnect()
	router := mux.NewRouter()
	ch := UserHandlers{service.NewUserService(domain.NewCassandraDb(DB))}
	routes(router, &ch)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
