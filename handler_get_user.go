package main

import (
	"fmt"
	"net/http"

	"github.com/aduatgit/rss-feed-aggregator/internal/auth"
)

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Couldn't get apikey", err))
		return
	}

	user, err := cfg.DB.GetUserByApi(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Couldn't retrieve user", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}
