package handlers

import "net/http"

// Login logs in the user, and adds a session to the database.
// Routed from POST "/login".
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

}

// Register validates the request and adds a new user to the database.
// Routed from POST "/register".
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

}
