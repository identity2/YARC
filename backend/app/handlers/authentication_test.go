package handlers

import (
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/login", h.Login)

	// Request body.
	validBody := `{"username":"mediocre_nothingness","password":"password"}`
	invalidBody := `{"username:}`
	nonExistUsernameBody := `{"username":"username","password":"ok"}`

	// Testcases.
	tests := []struct {
		name     string
		body     string
		wantCode int
	}{
		{"Valid Info", validBody, http.StatusOK},
		{"Invalid JSON", invalidBody, http.StatusBadRequest},
		{"Invalid Info", nonExistUsernameBody, http.StatusUnauthorized},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, _ := r.testHTTP(t, "POST", "/login", tc.body, "mediocrenotingness")

			if code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, code)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/register", h.Register)

	// Testcases.
	tests := []struct {
		name     string
		email    string
		wantCode int
		wantBody string
	}{
		{"Valid Info", "apple@google.com.tw", http.StatusCreated, ``},
		{"Duplicate", "mobile@blizzard.com.hk", http.StatusBadRequest, `{"error":"the email is already used"}`},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testHTTP(t, "POST", "/register", `{"username":"ok1","password":"ok1234","email":"`+tc.email+`"}`, "mediocre_nothingness")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want %s; got %s", tc.wantBody, body)
			}
		})
	}
}
