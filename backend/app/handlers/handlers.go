// Package handlers handles requests routed from the web app router.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

const (
	documentationURL = "https://github.com/YuChaoGithub/YARC/tree/master/backend"
)

// Handler contains a database so that all handlers could access it.
type Handler struct {
	Accounts   *models.AccountModel
	Subreddits *models.SubredditModel
	Comments   *models.CommentModel
}

// Home redirects the client to the GitHub documentation page.
// Routed from GET "/".
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, documentationURL, http.StatusSeeOther)
}

// Ping writes "ok" to the response body.
// The client would use this as a connection test.
// Routed from GET "/ping".
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
