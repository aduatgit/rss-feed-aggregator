package main

import (
	"fmt"
	"net/http"

	"github.com/aduatgit/rss-feed-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedidstr := chi.URLParam(r, "feedFollowID")
	feedid, err := uuid.Parse(feedidstr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}

	err = cfg.DB.DeleteFollow(r.Context(), database.DeleteFollowParams{
		ID:     feedid,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete follow on feed: %v", err))
		return
	}

	respondWithJson(w, 200, struct{}{})
}
