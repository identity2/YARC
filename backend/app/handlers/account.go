package handlers

import "net/http"

// User responds with the information of a user.
// Routed from GET "/user/{username}".
func (h *Handler) User(w http.ResponseWriter, r *http.Request) {

}

// ModifyBio changes the short bio of the current user.
// Routed from PUT "/me/bio".
func (h *Handler) ModifyBio(w http.ResponseWriter, r *http.Request) {

}

// SaveArticle saves the article for the current user.
// Routed from POST "/me/save/{articleID}".
func (h *Handler) SaveArticle(w http.ResponseWriter, r *http.Request) {

}

// JoinSubreddit adds the current user in the subscribe list of a subreddit.
func (h *Handler) JoinSubreddit(w http.ResponseWriter, r *http.Request) {

}
