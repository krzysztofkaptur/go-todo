package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	address string
	store   Database
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", makeHTTPHandleFunc(s.handleTodos))
	router.HandleFunc("/todos/{id}", makeHTTPHandleFunc(s.handleTodo))

	http.ListenAndServe(s.address, router)
}



type ApiError struct {
	Error string `json:"error"`
}

