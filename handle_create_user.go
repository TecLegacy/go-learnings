package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/teclegacy/rss-aggregator/internal/database"
)

func (apicfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decode := json.NewDecoder(r.Body)

	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to parse data %v", err))
		return
	}

	usr, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating User %v", err))
		return
	}

	responseWithJson(w, 201, dbUserToUser(usr))

}
