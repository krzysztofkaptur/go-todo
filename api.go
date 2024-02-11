package main

import (
	"net/http"
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
	router := http.NewServeMux()

	router.HandleFunc("GET /todos/", makeHTTPHandleFunc(s.handleFetchTodos))
	router.HandleFunc("POST /todos/", makeHTTPHandleFunc(s.handleCreateTodo))
	
	router.HandleFunc("GET /todos/{id}/", makeHTTPHandleFunc(s.handleFetchTodo))
	router.HandleFunc("DELETE /todos/{id}/", makeHTTPHandleFunc(s.handleDeleteTodo))
	router.HandleFunc("PATCH /todos/{id}/", makeHTTPHandleFunc(s.handleUpdateTodo))

	http.ListenAndServe(s.address, router)
}



type ApiError struct {
	Error string `json:"error"`
}

