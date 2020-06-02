// Package handlers handles requests routed from the web app router.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

const (
	documentationURL = "https://github.com/YuChaoGithub/YARC/tree/master/backend"
)

// Handler contains a database so that all handlers could access it.
type Handler struct {
	Accounts     *models.AccountModel
	Subreddits   *models.SubredditModel
	Comments     *models.CommentModel
	JWTSecretKey string
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

// jsonResponse responds with a JSON formatted payload.
func jsonResponse(w http.ResponseWriter, resp interface{}) {
	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write(json)
}

// respondWithError responds with an error code and a JSON formatted error message in the payload.
func respondWithError(w http.ResponseWriter, statusCode int, err error) {
	resp := struct {
		ErrStr string `json:"error"`
	}{
		ErrStr: err.Error(),
	}

	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json")
	w.Write(json)
}
