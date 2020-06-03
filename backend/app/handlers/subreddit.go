package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type subredditReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Subreddit responds with the information of a subreddit according to the subreddit name.
// Routed from GET "/subreddit/{name}".
func (h *Handler) Subreddit(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	resp, err := h.Subreddits.Get(name)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	jsonResponse(w, resp)
}

// NewSubreddit adds a new subreddit to the database.
// Routed from POST "/subreddit".
func (h *Handler) NewSubreddit(w http.ResponseWriter, r *http.Request) {
	// Retreive the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal json to subredditReq.
	req := subredditReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert to the database.
	err = h.Subreddits.Insert(req.Name, req.Description)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Successfully added new subreddit.
	w.WriteHeader(http.StatusCreated)
}

// TrendingSubreddit returns a list of trending subreddits.
// Routed from GET "/trending".
func (h *Handler) TrendingSubreddit(w http.ResponseWriter, r *http.Request) {
	// TODO
}
