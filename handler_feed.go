package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/teclegacy/rss-aggregator/internal/database"
)

func (apiCfg apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := &parameters{}

	//Serialize response body
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could'nt Parse User data %v", err))
		return
	}

	// Create Feed
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UsersID:   user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not Create Feed %v", err))
		return
	}

	responseWithJson(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	allFeeds, err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("No feeds found %v", err))
		return
	}

	// responseWithJson(w, 200, databaseFeedToFeed(allFeeds))
	log.Printf("Ss %v", allFeeds)
	// responseWithJson(w, 200, struct{}{})
	responseWithJson(w, 200, databaseALLFeedsToFeed(allFeeds))

}
