package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aduatgit/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type req struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	type resp struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}

	decoder := json.NewDecoder(r.Body)
	params := req{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error decoding JSON")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Couldn't create feed: ", err))
	}

	feedfollow, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error following feed: %v", err))
		return
	}

	response := resp{
		Feed:       databaseFeedToFeed(feed),
		FeedFollow: databaseFeedFollowToFeedFollow(feedfollow),
	}

	respondWithJson(w, 200, response)
}
