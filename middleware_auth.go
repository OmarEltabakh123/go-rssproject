package main

import (
	"fmt"
	"net/http"

	"github.com/OmarEltabakh123/rssproject/internal/auth"
	"github.com/OmarEltabakh123/rssproject/internal/database"
)

type authedhandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedhandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth erorr: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("could't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
