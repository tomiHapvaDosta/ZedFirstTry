package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tomiHapvaDosta/ZedFirstTry/internal/auth"
)

func (cfg *apiConfig) loginUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	requestParams := requestParamsUser{}
	if err := decoder.Decode(&requestParams); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := cfg.queries.GetUser(r.Context(), requestParams.Email)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	match, err := auth.CheckPasswordHash(requestParams.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !match {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.tokenSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{Token: token})
}
