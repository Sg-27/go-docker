package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHealthzHandler tests the /healtz endpoint
func TestHealthzHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healtz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthzHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "OK\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
