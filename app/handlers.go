package app

import (
	"cassandra_rest_api_users/domain"
	"cassandra_rest_api_users/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandlers struct {
	service service.UserService
}

func (h *UserHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *UserHandlers) Insert(w http.ResponseWriter, r *http.Request) {
	var user domain.UserStruct
	er1 := json.NewDecoder(r.Body).Decode(&user)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	result, er2 := h.service.Insert(&user)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *UserHandlers) GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "id cannot be empty", http.StatusBadRequest)
		return
	}
	result, er0 := h.service.GetById(id)
	if er0 != nil {
		http.Error(w, er0.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *UserHandlers) Update(w http.ResponseWriter, r *http.Request) {
	var user domain.UserStruct
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	result, er2 := h.service.Update(&user)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *UserHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	result, err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
