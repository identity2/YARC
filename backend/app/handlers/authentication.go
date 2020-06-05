package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Login logs in the user, and adds a session to the database.
// Routed from POST "/login".
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Retreive the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal json to logingReq.
	req := loginReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate the username and password.
	err = h.Accounts.Authenticate(req.Username, req.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}

	// Username and password validated, issue a JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"iat":      time.Now().Unix(),
	})

	tokenStr, err := token.SignedString([]byte(h.JWTSecretKey))
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the token in JSON the response body.
	resp := struct {
		Token string `json:"token"`
	}{
		Token: tokenStr,
	}
	jsonResponse(w, http.StatusOK, resp)
}

// Register validates the request and adds a new user to the database.
// Routed from POST "/register".
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// Retreive the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Unmarshal to json to registerReq.
	req := registerReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert into the database.
	err = h.Accounts.Insert(req.Username, req.Email, req.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Registered successfully.
	w.WriteHeader(http.StatusCreated)
}
