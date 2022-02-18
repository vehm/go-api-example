package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// Singular Todo
type Todo struct {
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

// Slice of Todos
type Todos []Todo

type todoHandler struct {
	sync.Mutex
	todos Todos
}

// Get All Todos
func (h *todoHandler) get(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.todos)
}

// Add Todo
func (h *todoHandler) post(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	var t Todo
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.todos = append(h.todos, t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.todos)
}

// ServeHTTP implements http.Handler, which means the todoHandler struct can be used in http.Handle
func (h *todoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
	case "POST":
		h.post(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// New Todo Handler
func newTodoHandler() *todoHandler {
	return &todoHandler{
		todos: Todos{},
	}
}
