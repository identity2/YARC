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

// Design notes: The point of defining all these interfaces is that
// we could insert stubs with the same interfaces to perform tests.

// AccountModel defines the operations related to account.
type AccountModel interface {
	Insert(string, string, string) error
	Authenticate(string, string) error
	ModifyBio(string, string) error
	SaveArticle(string, string) error
	JoinSubreddit(string, string) error
	Get(string) (models.UserInfo, error)
}

// SubredditModel defines the operations related to subreddit.
type SubredditModel interface {
	Insert(string, string) error
	Get(string) (models.SubredditInfo, error)
	GetTrending(int) ([]models.SubredditInfo, error)
}

// CommentModel defines the operations related to comment.
type CommentModel interface {
	Insert(string, string, string, string) (string, error)
	ModifyBody(string, string, string) error
	Delete(string, string) error
	GetByArticle(string, string, int) ([]models.CommentInfo, error)
	GetByUsername(string, string, int) ([]models.CommentInfo, error)
	GetPoints(string) (int, error)
}

// ArticleModel defines the operations related to article.
type ArticleModel interface {
	Insert(string, string, string, string, string) (string, error)
	ModifyBody(string, string, string) error
	Delete(string, string) error
	Get(string) (models.ArticleInfo, error)
	GetBySubreddit(string, string, string, int) ([]models.ArticleInfo, error)
	GetByUser(string, string, string, int) ([]models.ArticleInfo, error)
	GetSavedByUser(string, string, string, int) ([]models.ArticleInfo, error)
	GetPoints(string) (int, error)
}

// SearchModel defines the operations related to searching.
type SearchModel interface {
	GetResult(string) (models.SearchResults, error)
}

// Handler contains a database so that all handlers could access it.
type Handler struct {
	Accounts     AccountModel
	Subreddits   SubredditModel
	Comments     CommentModel
	Articles     ArticleModel
	Searches     SearchModel
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
func jsonResponse(w http.ResponseWriter, statusCode int, resp interface{}) {
	json, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
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

	json, _ := json.Marshal(resp)

	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json")
	w.Write(json)
}
