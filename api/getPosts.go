package main

import "net/http"

func (cfg *apiConfig) getPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := cfg.queries.GetPosts(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
