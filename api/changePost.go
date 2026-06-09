package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/tomiHapvaDosta/ZedFirstTry/internal/auth"
	"github.com/tomiHapvaDosta/ZedFirstTry/internal/database"
)

func (cfg *apiConfig) changePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	requestParamsPost := requestParamsPost{}
	if err := decoder.Decode(&requestParamsPost); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	postID := r.PathValue("id")
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	post1, err := cfg.queries.GetPost(r.Context(), postUUID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userUUID := post1.UserID.UUID

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	requestUserUUID, err := auth.ValidateJWT(token, cfg.tokenSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	if userUUID != requestUserUUID {
		respondWithError(w, http.StatusUnauthorized, "forbidden")
		return
	}

	post, err := cfg.queries.UpdatePost(r.Context(), database.UpdatePostParams{Title: requestParamsPost.Title, Body: requestParamsPost.Body, ID: postUUID})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, post)
}
