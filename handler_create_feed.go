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

	respondWithJson(w, 200, databaseFeedToFeed(feed))
}
