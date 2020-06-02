package handlers

import (
	"fmt"
	"net/http"
)

// ListArticle responds with a list of articles according the the query string.
// Routed from GET "/article".
func (h *Handler) ListArticle(w http.ResponseWriter, r *http.Request) {

}

// Article responds with an article according to the articleID parameter.
// Routed from GET "/article/{id}".
func (h *Handler) Article(w http.ResponseWriter, r *http.Request) {

}

// NewArticle adds a new article to the database.
// Routed from POST "/article".
func (h *Handler) NewArticle(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(usernameCtxKey)
	w.Write([]byte(fmt.Sprintf("The username is %v\n", username)))
}

// ModifyArticle changes an article to the new content.
// Routed from PUT "/article/{id}".
func (h *Handler) ModifyArticle(w http.ResponseWriter, r *http.Request) {

}

// DeleteArticle deletes an article in the database.
// Routed from DELETE "/article/{id}".
func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {

}
