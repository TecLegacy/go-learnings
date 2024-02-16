-- name: CreateFeed :one

INSERT INTO feeds(id,created_At,updated_At,name,url,users_id)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING *;


-- name: GetFeed :many
select * 
from feeds;

