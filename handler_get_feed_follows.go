package main

import (
	"net/http"

	"github.com/aduatgit/rss-feed-aggregator/internal/database"
)

func (cfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollows, err := cfg.DB.GetUserFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, "Couldn't get feed follows from Database")
	}
	respondWithJson(w, 200, databaseFeedFollowsToFeedFollows(feedfollows))
}
