// Package handlers handles requests routed from the web app router.
package handlers

import (
	"encoding/json"
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
	Insert(username, email, password string) error
	Authenticate(username, password string) error
	ModifyBio(username, newBio string) error
	SaveArticle(articleID, username string) error
	UnsaveArticle(articleID, username string) error
	JoinSubreddit(subName, username string) error
	LeaveSubreddit(subName, username string) error
	Get(username string) (models.UserInfo, error)
}

// SubredditModel defines the operations related to subreddit.
type SubredditModel interface {
	Insert(name, description string) error
	Get(name string) (models.SubredditInfo, error)
	GetTrending(limit int) ([]models.SubredditInfo, error)
}

// CommentModel defines the operations related to comment.
type CommentModel interface {
	Insert(subName, articleID, body, postedBy string) (string, error)
	ModifyBody(commentID, username, newBody string) error
	Delete(commentID, username string) error
	GetByArticle(articleID, afterCommentID string, limit int) ([]models.CommentInfo, error)
	GetByUsername(username, afterCommentID string, limit int) ([]models.CommentInfo, error)
}

// ArticleModel defines the operations related to article.
type ArticleModel interface {
	Insert(subName, postType, body, title, postedBy string) (string, error)
	ModifyBody(articleID, username, body string) error
	Delete(articleID, username string) error
	Get(articleID string) (models.ArticleInfo, error)
	GetBySubreddit(subName, afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error)
	GetByUser(username, afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error)
	GetSavedByUser(username, afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error)
	GetAll(afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error)
	GetBySubscribed(username, afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error)
}

// SearchModel defines the operations related to searching.
type SearchModel interface {
	GetResult(query string) (models.SearchResults, error)
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
	w.Write([]byte("ok"))
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
