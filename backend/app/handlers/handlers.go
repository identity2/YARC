// Package handlers handles requests routed from the web app router.
package handlers

import (
	"net/http"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

// Handler contains a database so that all handlers could access it.
type Handler struct {
	Accounts   *models.AccountModel
	Subreddits *models.SubredditModel
	Comments   *models.CommentModel
}

// Home simply writes "ok" to the response body to indicate
// that the connection to the server is ok.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
