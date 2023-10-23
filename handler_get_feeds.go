package main

import "net/http"

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, "Couldn't get feeds from Database")
	}
	respondWithJson(w, 200, databaseFeedsToFeeds(feeds))
}
