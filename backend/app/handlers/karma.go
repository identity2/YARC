package handlers

import "net/http"

// VoteArticle increases or decreases the points of an article.
// Routed from POST "/karma/article/{articleID}".
func (h *Handler) VoteArticle(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// VoteComment increase or decreases the points of a comment.
// Routed from POST "/karma/article/{commentID}".
func (h *Handler) VoteComment(w http.ResponseWriter, r *http.Request) {
	// TODO
}
