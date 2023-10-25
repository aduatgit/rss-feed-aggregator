package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aduatgit/rss-feed-aggregator/internal/database"
)

func (cfg *apiConfig) handlerGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limitString := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil || limitString == "" {
		limit = 10
	}
	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		ID:    user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't fetch posts from db: %v", err))
		return
	}
	respondWithJson(w, 200, databasePostsToPosts(posts))
}
