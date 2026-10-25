package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/OmarEltabakh123/rssproject/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body) // r.body >> request data from clients

	params := parameters{}
	err := decoder.Decode(&params) // Decode(&params) >> convert json to go
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{ //apiCfg.DB.CreateUsers >> save data in DB
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could't create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeed(r.Context()) //apiCfg.DB.CreateUsers >> save data in DB
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could't get feeds: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
