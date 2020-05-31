package handlers

import "net/http"

// Subreddit responds with the information of a subreddit according to the subreddit name.
// Routed from GET "/subreddit/{name}".
func (h *Handler) Subreddit(w http.ResponseWriter, r *http.Request) {

}

// NewSubreddit adds a new subreddit to the database.
// Routed from POST "/subreddit".
func (h *Handler) NewSubreddit(w http.ResponseWriter, r *http.Request) {

}

// TrendingSubreddit returns a list of trending subreddits.
// Routed from GET "/trending".
func (h *Handler) TrendingSubreddit(w http.ResponseWriter, r *http.Request) {

}
