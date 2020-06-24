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

// VoteArticle increases or decreases the points of an article.
// Routed from POST "/karma/article/{articleID}".
func (h *Handler) VoteArticle(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user.
	username := r.Context().Value(usernameCtxKey).(string)

	// Parameter.
	articleID := mux.Vars(r)["articleID"]

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
// Routed from POST "/karma/comment/{commentID}".
func (h *Handler) VoteComment(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user.
	username := r.Context().Value(usernameCtxKey).(string)

	// Parameter.
	commentID := mux.Vars(r)["commentID"]

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
