package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/tomiHapvaDosta/ZedFirstTry/internal/auth"
)

func (cfg *apiConfig) deletePost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	userUUIDEnd, err := auth.ValidateJWT(token, cfg.tokenSecret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

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

	userUUIDPost := post.UserID.UUID

	if userUUIDEnd != userUUIDPost {
		respondWithError(w, http.StatusUnauthorized, "forbidden")
		return
	}

	if err := cfg.queries.DeletePost(r.Context(), userUUIDPost); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
