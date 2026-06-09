package main

import (
	"encoding/json"
	"net/http"

	"github.com/tomiHapvaDosta/ZedFirstTry/internal/auth"
	"github.com/tomiHapvaDosta/ZedFirstTry/internal/database"
)

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	requestParams := requestParamsUser{}
	if err := decoder.Decode(&requestParams); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	hashedPassword, err := auth.HashPassword(requestParams.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := cfg.queries.CreateUser(r.Context(), database.CreateUserParams{Email: requestParams.Email, HashedPassword: hashedPassword})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respontWithJSON(w, http.StatusCreated, user)
}
