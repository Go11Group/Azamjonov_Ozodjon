package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Azamjonov_Ozodjon/lesson35/model"
	"github.com/Azamjonov_Ozodjon/lesson35/storage/postgres"
	"github.com/gorilla/mux"
	"net/http"
)

type Handleru struct {
	users postgres.UsersRepo
}

type Handlerp struct {
	problems postgres.ProblemsRepo
}

type Handlers struct {
	sproblems postgres.SolvedProblemsRepo
}

func NewHandler(users postgres.UsersRepo, problems postgres.ProblemsRepo, sproblems postgres.SolvedProblemsRepo) *http.Server {
	handleru := Handleru{users: users}
	handlerp := Handlerp{problems: problems}
	handlers := Handlers{sproblems: sproblems}
	m := mux.NewRouter()
	m.HandleFunc("/users", handleru.CreateUsers).Methods("POST")
	m.HandleFunc("/users/{id}", handleru.GetUsers).Methods("GET")
	m.HandleFunc("/users/{id}", handleru.UpdateUsers).Methods("PUT")
	m.HandleFunc("/users/{id}", handleru.DeleteUsers).Methods("DELETE")

	m.HandleFunc("/problems", handlerp.CreateProblems).Methods("POST")
	m.HandleFunc("/problems/{id}", handlerp.GetProblems).Methods("GET")
	m.HandleFunc("/problems/{id}", handlerp.UpdateProblems).Methods("PUT")
	m.HandleFunc("/problems/{id}", handlerp.DeleteProblems).Methods("DELETE")

	m.HandleFunc("/solved_problems", handlers.CreateSProblems).Methods("POST")
	m.HandleFunc("/solved_problems/{id}", handlers.GetSProblems).Methods("GET")
	m.HandleFunc("/solved_problems/{id}", handlers.UpdateSProblems).Methods("PUT")
	m.HandleFunc("/solved_problems/{id}", handlers.DeleteSProblems).Methods("DELETE")

	return &http.Server{Addr: ":8080", Handler: m}
}

func (h *Handleru) CreateUsers(w http.ResponseWriter, r *http.Request) {
	var user model.Users

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR!"))
		return
	}

	err = h.users.Create(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR!"))
		return
	}

	fmt.Println(user)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 OK OK OK!"))
}

func (p *Handlerp) CreateProblems(w http.ResponseWriter, r *http.Request) {
	var problems model.Problems

	err := json.NewDecoder(r.Body).Decode(&problems)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR!"))
		return
	}

	err = p.problems.Create(problems)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR!"))
		return
	}

	fmt.Println(problems)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 OK OK OK!"))
}

func (h *Handleru) GetUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.users.GetByID(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found!"))
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (p *Handlerp) GetProblems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := p.problems.GetByID(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Problem not found!"))
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handleru) UpdateUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user model.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}
	user.Id = id

	err = h.users.Update(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error updating user"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func (p *Handlerp) UpdateProblems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var problems model.Problems
	err := json.NewDecoder(r.Body).Decode(&problems)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}
	problems.Id = id

	err = p.problems.Update(problems)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error updating problem"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Problem updated successfully"))
}

func (h *Handleru) DeleteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.users.Delete(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error deleting user"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func (p *Handlerp) DeleteProblems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := p.problems.Delete(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error deleting problem"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Problem deleted successfully"))
}

func (sp *Handlers) CreateSProblems(w http.ResponseWriter, r *http.Request) {
	var sproblems model.SolvedProblem

	// Decode JSON request body into SolvedProblem struct
	err := json.NewDecoder(r.Body).Decode(&sproblems)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	// Validate the user_id and problem_id fields
	if sproblems.UserID == 0 || sproblems.ProblemID == 0 {
		http.Error(w, "User ID or Problem ID cannot be zero", http.StatusBadRequest)
		return
	}

	// Call the Create method of the SolvedProblemsRepo
	err = sp.sproblems.Create(sproblems)
	if err != nil {
		http.Error(w, "Failed to create solved problem", http.StatusInternalServerError)
		return
	}

	// Respond with success status
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Solved problem created successfully"))
}

func (sp *Handlers) GetSProblems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sproblems, err := sp.sproblems.GetByID(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Solved problem not found!"))
		return
	}

	json.NewEncoder(w).Encode(sproblems)
}

func (sp *Handlers) UpdateSProblems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var sproblems model.SolvedProblem
	err := json.NewDecoder(r.Body).Decode(&sproblems)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}
	sproblems.Id = id

	err = sp.sproblems.UpdateByID(id, sproblems)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error updating solved problem"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Solved problem updated successfully"))
}

func (sp *Handlers) DeleteSProblems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := sp.sproblems.DeleteByID(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error deleting solved problem"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Solved problem deleted successfully"))
}
