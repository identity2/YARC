package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User responds with the information of a user.
// Routed from GET "/user/{username}".
func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	resp, err := h.Accounts.Get(username)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	jsonResponse(w, http.StatusOK, resp)
}

// ModifyBio changes the short bio of the current user.
// Routed from PUT "/me/bio".
func (h *Handler) ModifyBio(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user.
	username := r.Context().Value(usernameCtxKey).(string)

	// Retreive the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal json.
	req := struct {
		Bio string `json:"bio"`
	}{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Modify the entry in database.
	err = h.Accounts.ModifyBio(username, req.Bio)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Modified successfully.
	w.WriteHeader(http.StatusNoContent)
}

// SaveArticle saves the article for the current user.
// Routed from POST "/me/save/{articleID}".
func (h *Handler) SaveArticle(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user and the articleID.
	username := r.Context().Value(usernameCtxKey).(string)
	articleID := mux.Vars(r)["articleID"]

	// Insert to database.
	err := h.Accounts.SaveArticle(articleID, username)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UnsaveArticle unsaves the article for the current user.
// Routed from POST /me/unsave/{articleID}
func (h *Handler) UnsaveArticle(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user and the articleID.
	username := r.Context().Value(usernameCtxKey).(string)
	articleID := mux.Vars(r)["articleID"]

	// Delete the entry in database.
	err := h.Accounts.UnsaveArticle(articleID, username)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	// Successfully deleted.
	w.WriteHeader(http.StatusNoContent)
}

// JoinState responds with the joining state of the user in the current subreddit.
func (h *Handler) JoinState(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user and the subName.
	username := r.Context().Value(usernameCtxKey).(string)
	subName := mux.Vars(r)["subreddit"]

	// Get the state from the database.
	resp := struct {
		Joined bool `json:"joined"`
	}{h.Accounts.GetJoinState(username, subName)}

	jsonResponse(w, http.StatusOK, resp)
}

// JoinSubreddit adds the current user in the subscribe list of a subreddit.
// Routed from POST /me/join/{subreddit}
func (h *Handler) JoinSubreddit(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user and the subName.
	username := r.Context().Value(usernameCtxKey).(string)
	subName := mux.Vars(r)["subreddit"]

	// Insert to database.
	err := h.Accounts.JoinSubreddit(subName, username)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	// Increment the subreddit visit count.
	h.Subreddits.IncrVisitCount(subName)

	// Successfully inserted.
	w.WriteHeader(http.StatusCreated)
}

// LeaveSubreddit removes the current user from the subscribe list of a subreddit.
// Routed from POST /me/leave/{subreddit}
func (h *Handler) LeaveSubreddit(w http.ResponseWriter, r *http.Request) {
	// Retreive the currently logged in user and the subName.
	username := r.Context().Value(usernameCtxKey).(string)
	subName := mux.Vars(r)["subreddit"]

	// Delete the entry in database.
	err := h.Accounts.LeaveSubreddit(subName, username)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	// Successfully deleted.
	w.WriteHeader(http.StatusNoContent)
}
