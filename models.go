package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/teclegacy/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAt time.Time `json:"updated_At"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func dbUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAt time.Time `json:"updated_At"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UsersID   uuid.UUID `json:"userId"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UsersID:   dbFeed.UsersID,
	}
}
