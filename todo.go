package main

import (
	"encoding/json"
	"go-todo/internal/database"
	"net/http"
	"strconv"
)

type CreateTodoReq struct {
	Text string `json: "text"` 
}

type UpdateTodoReq struct {
	Text string `json: "text"`
}

func (s *ApiServer) handleFetchTodos(w http.ResponseWriter, r *http.Request) error {
	todos, err := s.store.DB.FetchTodos(r.Context())

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todos)
}

func (s *ApiServer) handleFetchTodo(w http.ResponseWriter, r *http.Request) error {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		return err
	}

	todo, err := s.store.DB.FetchTodo(r.Context(), int32(id))

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

func (s *ApiServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) error {
	createTodoReq := &CreateTodoReq{}

	err := json.NewDecoder(r.Body).Decode(&createTodoReq)

	if err != nil {
		return err
	}

	todo, err := s.store.DB.CreateTodo(r.Context(), createTodoReq.Text)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

func (s *ApiServer) handleDeleteTodo(w http.ResponseWriter, r *http.Request) error {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		return err
	}

	todo, err := s.store.DB.DeleteTodo(r.Context(), int32(id))

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

func (s *ApiServer) handleUpdateTodo(w http.ResponseWriter, r *http.Request) error {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	updateTodoReq := &UpdateTodoReq{}

	if err != nil {
		return err
	}

	json.NewDecoder(r.Body).Decode(&updateTodoReq)

	todo, err := s.store.DB.UpdateTodo(r.Context(), database.UpdateTodoParams{ID: int32(id), Text: updateTodoReq.Text})

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}