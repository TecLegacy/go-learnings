-- +goose Up 
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_At TIMESTAMP NOT NULL,
    updated_At TIMESTAMP NOT NULL,
    users_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feeds_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down 
DROP TABLE feed_follows;
