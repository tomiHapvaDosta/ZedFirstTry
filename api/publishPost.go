package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) publishPost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = cfg.queries.PublishPost(r.Context(), postUUID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return

	}

	client := http.Client{}
	body := struct {
		Event  string    `json:"event"`
		PostID uuid.UUID `json:"post_id"`
	}{Event: "post.published", PostID: postUUID}
	marshaledBody, err := json.Marshal(body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	wrappedBody := bytes.NewReader(marshaledBody)
	req, err := http.NewRequest("POST", "http://localhost:8080/api/webhooks", wrappedBody)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := cfg.queries.GetPost(r.Context(), postUUID)

	respondWithJSON(w, http.StatusOK, post)
}
