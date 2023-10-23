package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aduatgit/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type req struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := req{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedfollow, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error following feed: %v", err))
		return
	}

	respondWithJson(w, 200, databaseFeedFollowToFeedFollow(feedfollow))
}
