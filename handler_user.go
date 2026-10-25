package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/OmarEltabakh123/rssproject/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body) // r.body >> request data from clients

	params := parameters{}
	err := decoder.Decode(&params) // Decode(&params) >> convert json to go
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUsers(r.Context(), database.CreateUsersParams{ //apiCfg.DB.CreateUsers >> save data in DB
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could't get posts: %v", err))
		return
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
