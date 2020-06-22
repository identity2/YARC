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

// newCommentReq is the request body for posting a new comment.
type newCommentReq struct {
	ArticleID string `json:"articleID"`
	Body      string `json:"body"`
}

// ListComment returns a list of comments according to the query strings.
// Routed from GET "/comment".
func (h *Handler) ListComment(w http.ResponseWriter, r *http.Request) {
	// Get query parameters.
	q := r.URL.Query()
	after, ofArticle, postedBy := q.Get("after"), q.Get("ofArticle"), q.Get("postedBy")

	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		limit = defaultListLimit
	}

	// The response is a list of CommentInfo.
	resp := struct {
		Comments []models.CommentInfo `json:"comments"`
	}{Comments: []models.CommentInfo{}}

	// Check which criterion to use.
	if ofArticle != "" {
		resp.Comments, err = h.Comments.GetByArticle(ofArticle, after, limit)
	} else if postedBy != "" {
		resp.Comments, err = h.Comments.GetByUsername(postedBy, after, limit)
	} else {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("either the 'ofArticle' or 'postedBy' query string should be in the requested URL"))
		return
	}
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Successfully get the comments.
	jsonResponse(w, http.StatusOK, resp)
}

// Comment returns a comment according to the commentID.
// Routed from GET "/comment/{id}".
func (h *Handler) Comment(w http.ResponseWriter, r *http.Request) {
	commentID := mux.Vars(r)["id"]
	resp, err := h.Comments.Get(commentID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	jsonResponse(w, http.StatusOK, resp)
}

// NewComment adds a new comment to the database.
// Routed from POST "/comment".
func (h *Handler) NewComment(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user.
	username := r.Context().Value(usernameCtxKey).(string)

	// Retreive the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal json to newCommentReq.
	req := newCommentReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert to database.
	commentID, err := h.Comments.Insert(req.ArticleID, req.Body, username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Respond with the inserted commentID.
	resp := struct {
		CommentID string `json:"commentID"`
	}{commentID}

	jsonResponse(w, http.StatusCreated, resp)
}

// ModifyComment changes a comment to a new content.
// Routed from PUT "/comment/{id}".
func (h *Handler) ModifyComment(w http.ResponseWriter, r *http.Request) {
	// Retreive the current user logged in and the comment ID.
	username := r.Context().Value(usernameCtxKey).(string)
	commentID := mux.Vars(r)["id"]

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
	err = h.Comments.ModifyBody(commentID, username, req.Body)
	if err != nil {
		if errors.Is(err, models.ErrCommentNotExist) {
			respondWithError(w, http.StatusNotFound, err)
		} else {
			respondWithError(w, http.StatusBadRequest, err)
		}
		return
	}

	// Successfully modified.
	w.WriteHeader(http.StatusNoContent)
}

// DeleteComment deletes a comment in the database.
// Routed from DELETE "/comment/{id}".
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	// Retreive the current user logged in and the comment ID.
	username := r.Context().Value(usernameCtxKey).(string)
	commentID := mux.Vars(r)["id"]

	// Delete the entry in database.
	err := h.Comments.Delete(commentID, username)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	// Successfully deleted.
	w.WriteHeader(http.StatusNoContent)
}
