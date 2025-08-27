package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func makeAdminToken() string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"uid": "00000000-0000-0000-0000-000000000001", "role": "admin", "exp": time.Now().Add(time.Hour).Unix()})
	s, _ := t.SignedString([]byte("test-secret"))
	return s
}

func TestShowroomCRUD(t *testing.T) {
	r, _ := setupTest(t)

	token := makeAdminToken()

	// create
	payload := map[string]any{"name": "Room A", "description": "desc", "capacity": 10}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/showrooms", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("create expected 201, got %d: %s", w.Code, w.Body.String())
	}

	// list
	req = httptest.NewRequest(http.MethodGet, "/api/showrooms", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("list expected 200, got %d", w.Code)
	}
}
