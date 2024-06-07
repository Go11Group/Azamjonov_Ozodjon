package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"users_and_produncts/model"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := model.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error with getting json from body: %s", err)))
		return
	}
	err = h.Users.CreateUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with adding user to database: %s", err)))
		return
	}
	w.Write([]byte("Yangi User Muvaffaqiyatli qo'shildi"))
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := model.Filter{}
	if query.Has("id") {
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error with getting Id filter: %s", err)))
			return
		}
		filter.Id = &id
	}
	if query.Has("username") {
		username := query.Get("username")
		filter.Username = &username
	}

	users, err := h.Users.GetUser(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with getting information from database: %s", err)))
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with converting information into json: %s", err)))
		return
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userToUpdate := model.User{}
	err := json.NewDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error with getting json from body: %s", err)))
		return
	}

	err = h.Users.UpdateUser(&userToUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error with updating user in database: %s", err)))
		return
	}
	w.Write([]byte("User muvaffaqiyatli o'zgartirildi "))
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("id") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: There is no ID in query"))
		return
	}
	id, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: ID is not in correct type"))
		return
	}
	err = h.Users.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with deleteing user from database: %s", err)))
		return
	}
	w.Write([]byte("User muvaffaqiyatli o'chirildi"))
}
