package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/teclegacy/rss-aggregator/internal/database"
)

// HTTP route /feeds_follow
// POST , authenticated Route (API_KEY)
func (apiCfg *apiConfig) handleFeedFollowsRequest(w http.ResponseWriter, r *http.Request, currentUser database.User) {

	type RequestParameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	var requestParams RequestParameters
	jsonDecoder := json.NewDecoder(r.Body)

	err := jsonDecoder.Decode(&requestParams)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to parse request data: %v", err))
		return
	}

	// Create a new feed follow in the database.
	newFeedFollow, err := apiCfg.DB.CreateFeedsFollows(r.Context(), database.CreateFeedsFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UsersID:   currentUser.ID,
		FeedsID:   requestParams.FeedId,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to create feed follow: %v", err))
		return
	}

	respondWithJson(w, 201, dbFeedFollowsToFeedFollows(newFeedFollow))
}
