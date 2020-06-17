package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/YuChaoGithub/YARC/backend/app/models"
	"github.com/gorilla/mux"
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
	// TODO
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
	}{
		ArticleID: articleID,
	}

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
