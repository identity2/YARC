package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/YuChaoGithub/YARC/backend/app/models"
	"github.com/gorilla/mux"
)

const (
	defaultTrendingLimit = 5
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
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	// Increment the subreddit visit count.
	h.Subreddits.IncrVisitCount(resp.Name)

	jsonResponse(w, http.StatusOK, resp)
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
	// Retrieve the limit query.
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = defaultTrendingLimit
	}

	// Get the trending subreddits.
	subs, err := h.Subreddits.GetTrending(limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := struct {
		Subreddits []models.SubredditInfo `json:"subreddits"`
	}{subs}

	// Successfully get the subreddit list.
	jsonResponse(w, http.StatusOK, res)
}
