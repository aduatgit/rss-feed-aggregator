package main

import (
	"net/http"

	"github.com/aduatgit/rss-feed-aggregator/internal/database"
)

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, databaseUserToUser(user))
}
