-- name: CreateFeedsFollows :one
INSERT INTO feed_follows(id,created_At,updated_At,users_id,feeds_id)
VALUES($1,$2,$3,$4,$5)
RETURNING * ;

