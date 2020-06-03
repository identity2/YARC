package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YuChaoGithub/YARC/backend/app/models/mock"
	"github.com/gorilla/mux"
)

type testRouter struct {
	*mux.Router
}

// newTestHandler sets up a handler with mock models.
func newTestHandler(t *testing.T) *Handler {
	return &Handler{
		Accounts:     &mock.AccountModel{},
		Subreddits:   &mock.SubredditModel{},
		Comments:     &mock.CommentModel{},
		Articles:     &mock.ArticleModel{},
		Searches:     &mock.SearchModel{},
		JWTSecretKey: "Why should I stay here?",
	}
}

// newTestRouter sets up a router with the given route.
func newTestRouter(t *testing.T, method, path string, f func(http.ResponseWriter, *http.Request)) *testRouter {
	router := mux.NewRouter()
	router.HandleFunc(path, f).Methods(method)
	return &testRouter{router}
}

// testGet sends a GET API request and returns the status code and response body.
func (router *testRouter) testGet(t *testing.T, urlPath string) (int, string) {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("GET", urlPath, nil))
	return w.Code, w.Body.String()
}
