package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/YuChaoGithub/YARC/backend/app/models"
	"github.com/gorilla/mux"
)

const (
	defaultListSort  = "hot"
	defaultListLimit = 25
)

type newArticleReq struct {
	Subreddit string `json:"subreddit"`
	Type      string `json:"type"`
	Body      string `json:"body"`
	Title     string `json:"title"`
}

// ListArticle responds with a list of articles according the the query string.
// Routed from GET "/article".
func (h *Handler) ListArticle(w http.ResponseWriter, r *http.Request) {
	// Get query parameters.
	q := r.URL.Query()
	sortedBy, after, criterion, key := q.Get("sort"), q.Get("after"), q.Get("criterion"), q.Get("key")
	if sortedBy == "" {
		sortedBy = defaultListSort
	}

	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		limit = defaultListLimit
	}

	// The response is a list of ArticleInfo.
	resp := struct {
		Articles []models.ArticleInfo `json:"articles"`
	}{Articles: []models.ArticleInfo{}}

	// Criterion or key is invalid: list articles from all subreddits or from all subscribed
	// subreddits if the user is logged in.
	if criterion == "" || key == "" {
		// Check if the user is logged in.
		if authHeader, ok := r.Header["Authorization"]; ok && len(authHeader[0]) > 10 {
			// This block inside the if-condition is not included in the unit tests,
			// it is only manually tested, so modify with care...need to find a way to test it.
			username, err := decodeJWT(authHeader[0][7:], h.JWTSecretKey)
			if err == nil {
				// The user is logged in.
				resp.Articles, err = h.Articles.GetBySubscribed(username, after, sortedBy, limit)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err)
					return
				}

				// Successfully get the article lists, if the list is empty, treat the user as not logged in.
				if after != "" || len(resp.Articles) != 0 {
					jsonResponse(w, http.StatusOK, resp)
					return
				}
			}
		}

		// The user is not logged in.
		resp.Articles, err = h.Articles.GetAll(after, sortedBy, limit)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		// Successfully get the article lists.
		jsonResponse(w, http.StatusOK, resp)
		return
	}

	// Get the functions for each criterion.
	listFuncs := map[string]func(string, string, string, int) ([]models.ArticleInfo, error){
		"sub":     h.Articles.GetBySubreddit,
		"by":      h.Articles.GetByUser,
		"savedBy": h.Articles.GetSavedByUser,
	}
	fn, ok := listFuncs[criterion]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("the criterion query string should be either 'sub', 'by', or 'savedBy'"))
		return
	}

	// Get the article lists from the database.
	resp.Articles, err = fn(key, after, sortedBy, limit)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Successfully get the articles.
	jsonResponse(w, http.StatusOK, resp)
}

// Article responds with an article according to the articleID parameter.
// Routed from GET "/article/{id}".
func (h *Handler) Article(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["id"]
	resp, err := h.Articles.Get(articleID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	// Increment the subreddit visit count.
	h.Subreddits.IncrVisitCount(resp.Subreddit)

	jsonResponse(w, http.StatusOK, resp)
}

// NewArticle adds a new article to the database.
// Routed from POST "/article".
func (h *Handler) NewArticle(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user.
	username := r.Context().Value(usernameCtxKey).(string)

	// Retreive the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal json to newArticleReq.
	req := newArticleReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert to database.
	articleID, err := h.Articles.Insert(req.Subreddit, req.Type, req.Body, req.Title, username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Respond with the inserted articleID.
	resp := struct {
		ArticleID string `json:"articleID"`
	}{articleID}

	jsonResponse(w, http.StatusCreated, resp)
}

// ModifyArticle changes an article to the new content.
// Routed from PUT "/article/{id}".
func (h *Handler) ModifyArticle(w http.ResponseWriter, r *http.Request) {
	// Retreive the current user logged in and the article ID.
	username := r.Context().Value(usernameCtxKey).(string)
	articleID := mux.Vars(r)["id"]

	// Retreive the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal json of request body.
	req := struct {
		Body string `json:"body"`
	}{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Modify the entry in database.
	err = h.Articles.ModifyBody(articleID, username, req.Body)
	if err != nil {
		if errors.Is(err, models.ErrArticleNotExist) {
			respondWithError(w, http.StatusNotFound, err)
		} else {
			respondWithError(w, http.StatusBadRequest, err)
		}
		return
	}

	// Successfully modified.
	w.WriteHeader(http.StatusNoContent)
}

// DeleteArticle deletes an article in the database.
// Routed from DELETE "/article/{id}".
func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// Retreive the current user logged in and the articleID.
	username := r.Context().Value(usernameCtxKey).(string)
	articleID := mux.Vars(r)["id"]

	// Delete the entry in database.
	err := h.Articles.Delete(articleID, username)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	// Successfully deleted.
	w.WriteHeader(http.StatusNoContent)
}
