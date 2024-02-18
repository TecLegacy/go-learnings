package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/teclegacy/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {

	type RequestParameters struct {
		Name string `json:"name"`
	}
	var requestParams RequestParameters

	jsonDecoder := json.NewDecoder(r.Body)

	err := jsonDecoder.Decode(&requestParams)
	if err != nil {

		respondWithError(w, 400, fmt.Sprintf("Failed to parse request data: %v", err))
		return
	}

	// Create a new user in the database.
	newUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      requestParams.Name,
	})
	if err != nil {

		respondWithError(w, 400, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

	respondWithJson(w, 201, dbUserToUser(newUser))
}

func (apiCfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, 200, dbUserToUser(user))
}
