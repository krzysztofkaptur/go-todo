package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id   int    `json: "id"`
	Text string `json: "text"`
}

type CreateTodoReq struct {
	Text string `json: "text"` 
}

type UpdateTodoReq struct {
	Text string `json: "text"`
}

func (s *ApiServer) handleTodos(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleFetchTodos(w, r)
	} else if r.Method == "POST" {
		return s.handleCreateTodo(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *ApiServer) handleTodo(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "DELETE" {
		return s.handleDeleteTodo(w, r)
	} else if r.Method == "GET" {
		return s.handleFetchTodo(w, r)
	} else if r.Method == "PATCH" {
		return s.handleUpdateTodo(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *ApiServer) handleFetchTodos(w http.ResponseWriter, r *http.Request) error {
	query := `
		select * 
		from todos;
	`
	rows, err := s.store.db.Query(query)

	if err != nil {
		return err
	}

	todos := []*Todo{}

	for rows.Next() {
		todo := &Todo{}
		
		err := rows.Scan(
			&todo.Id,
			&todo.Text,
		)

		if err != nil {
			return err
		}

		todos = append(todos, todo)
	}

	return WriteJSON(w, http.StatusOK, todos)
}

func (s *ApiServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) error {
	createTodoReq := &CreateTodoReq{}
	
	query := `
		insert into todos (text)
		values ($1);
	`

	err := json.NewDecoder(r.Body).Decode(&createTodoReq);

	if err != nil {
		return err
	}

	rows, err := s.store.db.Query(query, createTodoReq.Text)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, rows)
}

func (s *ApiServer) handleDeleteTodo(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	query := `
		delete from todos
		where id=$1;
	`

	_, err := s.store.db.Query(query, id)

	if err != nil {
		return err
	}

	// add some message that removal succeeded

	return nil
}

func (s *ApiServer) handleFetchTodo(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]


	query := `
		select *
		from todos
		where id=$1;
	`

	rows, err := s.store.db.Query(query, id)

	if err != nil {
		return err
	}

	todo := &Todo{}

	for rows.Next() {
		err := rows.Scan(
			&todo.Id,
			&todo.Text,
		)

		if err != nil {
			return err
		}
	}

	return WriteJSON(w, http.StatusOK, todo)
}

func (s *ApiServer) handleUpdateTodo(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	updateTodoReq := &UpdateTodoReq{}

	jsonErr := json.NewDecoder(r.Body).Decode(&updateTodoReq)

	if jsonErr != nil {
		return jsonErr
	}

	query := `
		update todos
		set text=$2
		where id=$1
	`

	_, dbErr := s.store.db.Query(query, id, updateTodoReq.Text)

	if dbErr != nil {
		return dbErr
	}

	return nil
}