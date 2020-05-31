package handlers

import "net/http"

// ListComment returns a list of comments according to the query strings.
// Routed from GET "/comment".
func (h *Handler) ListComment(w http.ResponseWriter, r *http.Request) {

}

// Comment returns a comment according to the commentID.
// Routed from GET "/comment/{id}".
func (h *Handler) Comment(w http.ResponseWriter, r *http.Request) {

}

// NewComment adds a new comment to the database.
// Routed from POST "/comment".
func (h *Handler) NewComment(w http.ResponseWriter, r *http.Request) {

}

// ModifyComment changes a comment to a new content.
// Routed from PUT "/comment/{id}".
func (h *Handler) ModifyComment(w http.ResponseWriter, r *http.Request) {

}

// DeleteComment deletes a comment in the database.
// Routed from DELETE "/comment/{id}".
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

}
