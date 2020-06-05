package handlers

import (
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/ping", h.Ping)

	// Perform test.
	t.Run("ping", func(t *testing.T) {
		code, body := r.testGet(t, "/ping")

		if code != http.StatusOK {
			t.Errorf("want 200; got %d", code)
		}

		if body != "ok" {
			t.Errorf("want ok; got %s", body)
		}
	})
}

func TestJSONResponse(t *testing.T) {
	type testStruct struct {
		Content int `json:"content"`
	}

	// Testcases.
	tests := []struct {
		name     string
		code     int
		resp     interface{}
		wantCode int
		wantBody string
	}{
		{"Valid", http.StatusOK, testStruct{3}, http.StatusOK, `{"content":3}`},
		{"Invalid", http.StatusOK, math.Inf(1), http.StatusInternalServerError, ``},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub.
			w := httptest.NewRecorder()

			// When.
			jsonResponse(w, tc.code, tc.resp)

			// Want.
			if w.Code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, w.Code)
			}

			body := w.Body.String()
			if body != tc.wantBody {
				t.Errorf("want %s; got %s", tc.wantBody, body)
			}
		})
	}
}

func TestRespondWithError(t *testing.T) {
	errTestError := errors.New("this is a test error")
	wantBody := `{"error":"` + errTestError.Error() + `"}`
	wantCode := http.StatusConflict

	t.Run("error response", func(t *testing.T) {
		// Stub.
		w := httptest.NewRecorder()

		// When.
		respondWithError(w, wantCode, errTestError)

		// Want.
		if w.Code != wantCode {
			t.Errorf("want %d; got %d", wantCode, w.Code)
		}

		body := w.Body.String()
		if body != wantBody {
			t.Errorf("want %s; got %s", wantBody, body)
		}
	})
}
