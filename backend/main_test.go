package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/ethirajmudhaliar/backend/react-go-gov-search-data/govData/v1"
	"github.com/gorilla/mux"
)

func TestLoggingMiddleware(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/v1/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	router.Use(LoggingMiddleware)

	req, err := http.NewRequest("GET", "/v1/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}
}

func TestRoutes(t *testing.T) {
	router := mux.NewRouter()

	router.HandleFunc("/api/data", v1.GetGovernmentData).Methods("GET")

	req, err := http.NewRequest("GET", "/api/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusNotFound, status)
	}
}

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()

	tests := []struct {
		method      string
		url         string
		body        []byte
		statusCode  int
		description string
	}{
		{"GET", "/receipts/non-existent-id/points", nil, http.StatusNotFound, "Get receipt points for non-existent receipt"},
	}

	for _, tt := range tests {
		req, err := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(tt.body))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != tt.statusCode {
			t.Errorf("%s: expected status code %d, got %d", tt.description, tt.statusCode, status)
		}
	}
}
