-- name: FetchTodos :many
SELECT * FROM todos;

-- name: FetchTodo :one
SELECT * FROM todos WHERE id=$1;

-- name: CreateTodo :one
INSERT INTO todos ("text") VALUES($1) RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todos WHERE id=$1 RETURNING *;

-- name: UpdateTodo :one
UPDATE todos SET text=COALESCE($2, text) WHERE id=$1 RETURNING *;