// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedsFollows = `-- name: CreateFeedsFollows :one
INSERT INTO feed_follows(id,created_At,updated_At,users_id,feeds_id)
VALUES($1,$2,$3,$4,$5)
RETURNING id, created_at, updated_at, users_id, feeds_id
`

type CreateFeedsFollowsParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UsersID   uuid.UUID
	FeedsID   uuid.UUID
}

func (q *Queries) CreateFeedsFollows(ctx context.Context, arg CreateFeedsFollowsParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedsFollows,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UsersID,
		arg.FeedsID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UsersID,
		&i.FeedsID,
	)
	return i, err
}
