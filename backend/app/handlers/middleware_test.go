package handlers

import (
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestDecodeJWT(t *testing.T) {
	secretKey := "I want you to know"
	username := "mediocre_nothingness"

	validTokenStr, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "iat": time.Now().Unix()}).SignedString([]byte(secretKey))
	expiredTokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1lZGlvY3JlX25vdGhpbmduZXNzIiwiaWF0IjoxMjU3ODk0MDAwfQ.yS5DR_di7d8xhx7rnPmCCdgwkQ-6hCiW8pNFUwdKSt0"
	invalidTokenStr := "xyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1lZGlvY3JlX25vdGhpbmduZXNzIiwiaWF0IjoxNTkwNDI2MTA4MDB9.JIkHdK42paiRLhbl6R7_dc7vOgDmKQnuVyEpaV275y0"

	// Testcases.
	tests := []struct {
		name         string
		tokenStr     string
		wantUsername string
		wantErrorStr string
	}{
		{"Valid Token", validTokenStr, username, ""},
		{"Expired Token", expiredTokenStr, "", "expired"},
		{"Invalid Token", invalidTokenStr, "", "validate"},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			username, err := decodeJWT(tc.tokenStr, secretKey)

			// Want.
			if username != tc.wantUsername {
				t.Errorf("want %s; got %s", tc.wantUsername, username)
			}

			if err != nil && !strings.Contains(err.Error(), tc.wantErrorStr) {
				t.Errorf("want error to contain %s; got %s", tc.wantErrorStr, err.Error())
			}
		})
	}

}
