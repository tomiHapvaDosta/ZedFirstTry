package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getPost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := cfg.queries.GetPost(r.Context(), postUUID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, post)
}
