package handlers

import (
	"fmt"
	"net/http"
)

// Search responds with the search result of user's query.
// Routed from GET "/search".
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	// Get the search query.
	q := r.URL.Query().Get("q")
	if q == "" {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("the query string cannot be empty"))
		return
	}

	// Search from the database.
	res, err := h.Searches.GetResult(q)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Successfully get the result.
	jsonResponse(w, http.StatusOK, res)
}
