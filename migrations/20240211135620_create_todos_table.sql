-- +goose Up
CREATE TABLE todos (id SERIAL PRIMARY KEY, text VARCHAR(50) NOT NULL);

-- +goose Down
DROP TABLE todos;

