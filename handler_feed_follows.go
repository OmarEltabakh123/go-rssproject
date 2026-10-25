package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/OmarEltabakh123/rssproject/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body) // r.body >> request data from clients

	params := parameters{}
	err := decoder.Decode(&params) // Decode(&params) >> convert json to go
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{ //apiCfg.DB.CreateUsers >> save data in DB
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could't create feed Follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID) //apiCfg.DB.CreateUsers >> save data in DB)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could't get feed Follows: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseGetFeedsToFeeds(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could't Parse feed Follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could't delete feed Follow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
