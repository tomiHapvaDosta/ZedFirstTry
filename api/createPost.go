package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/tomiHapvaDosta/ZedFirstTry/internal/auth"
	"github.com/tomiHapvaDosta/ZedFirstTry/internal/database"
)

func (cfg *apiConfig) createPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	requestParams := requestParamsPost{}
	if err := decoder.Decode(&requestParams); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tokenStr, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userUUID, err := auth.ValidateJWT(tokenStr, cfg.tokenSecret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	userNullUUID := uuid.NullUUID{UUID: userUUID, Valid: true}

	post, err := cfg.queries.CreatePost(r.Context(),
		database.CreatePostParams{UserID: userNullUUID, Title: requestParams.Title, Body: requestParams.Body,
			Published: sql.NullBool{Bool: false, Valid: true}})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respontWithJSON(w, http.StatusOK, post)
}
