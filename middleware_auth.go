package main

import (
	"fmt"
	"net/http"

	"github.com/aduatgit/rss-feed-aggregator/internal/auth"
	"github.com/aduatgit/rss-feed-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		handler(w, r, user)
	}
}
