-- +goose Up
CREATE TABLE users(
  id UUID PRIMARY KEY,
  created_At TIMESTAMP NOT NULL,
  updated_At TIMESTAMP NOT NULL,
  name Text NOT NULL
);

-- +goose Down 
DROP TABLE users;