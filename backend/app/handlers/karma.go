package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// points define the integer point according to each action.
var points = map[string]int{
	"up":     1,
	"down":   -1,
	"cancel": 0,
}

// GetArticleVote returns the voting action of the currently logged in user.
// Routed from GET "/me/karma/article/{id}"
func (h *Handler) GetArticleVote(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user.
	username := r.Context().Value(usernameCtxKey).(string)

	// Parameter.
	articleID := mux.Vars(r)["id"]

	// Retreive from database.
	point := h.Articles.GetVote(username, articleID)

	// Get the action string.
	action := "cancel"
	for key, val := range points {
		if val == point {
			action = key
			break
		}
	}

	// Respond to the user.
	resp := struct {
		Action string `json:"action"`
	}{action}

	jsonResponse(w, http.StatusOK, resp)
}

// GetCommentVote returns the voting action of the currently logged in user.
// Routed from GET "/me/karma/comment/{id}"
func (h *Handler) GetCommentVote(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user.
	username := r.Context().Value(usernameCtxKey).(string)

	// Parameter.
	commentID := mux.Vars(r)["id"]

	// Retreive from database.
	point := h.Comments.GetVote(username, commentID)

	// Get the action string.
	action := "cancel"
	for key, val := range points {
		if val == point {
			action = key
			break
		}
	}

	// Respond to the user.
	resp := struct {
		Action string `json:"action"`
	}{action}

	jsonResponse(w, http.StatusOK, resp)
}

// VoteArticle increases or decreases the points of an article.
// Routed from POST "/karma/article/{id}".
func (h *Handler) VoteArticle(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user.
	username := r.Context().Value(usernameCtxKey).(string)

	// Parameter.
	articleID := mux.Vars(r)["id"]

	// Query string.
	action := r.URL.Query().Get("action")
	point, ok := points[action]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("the action query string is required"))
		return
	}

	// Update the database.
	err := h.Articles.Vote(username, articleID, point)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Successfully updated the database.
	w.WriteHeader(http.StatusNoContent)
}

// VoteComment increase or decreases the points of a comment.
// Routed from POST "/karma/comment/{id}".
func (h *Handler) VoteComment(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user.
	username := r.Context().Value(usernameCtxKey).(string)

	// Parameter.
	commentID := mux.Vars(r)["id"]

	// Query string.
	action := r.URL.Query().Get("action")
	point, ok := points[action]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("the action query string is required"))
		return
	}

	// Update the database.
	err := h.Comments.Vote(username, commentID, point)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Successfully updated the database.
	w.WriteHeader(http.StatusNoContent)
}
