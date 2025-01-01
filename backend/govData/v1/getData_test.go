package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetGovernmentDataSuccess tests the GetGovernmentData function with a mock API response.
func TestGetGovernmentDataSuccess(t *testing.T) {
	// Mock the external API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
            ["NAME", "B01001_001E", "state"],
            ["Alabama", "5024279", "01"]
        ]`))
	}))
	defer mockServer.Close()

	// Replace the default HTTP client with a custom client that redirects requests to the mock server
	http.DefaultTransport = &mockTransport{
		base:     http.DefaultTransport,
		mockHost: mockServer.URL,
	}

	// Create a test request
	req, err := http.NewRequest("GET", "/api/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetGovernmentData)
	handler.ServeHTTP(rr, req)

	// Assert the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	// Assert the response body contains valid JSON with a "success" field
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}

	// Assert response structure
	if response["success"] != true {
		t.Errorf("expected success to be true, got %v", response["success"])
	}

	if _, ok := response["data"]; !ok {
		t.Fatalf("response does not contain 'data' field")
	}

	// Check that the "data" field is an array and contains one item
	data, ok := response["data"].([]interface{})
	if !ok {
		t.Fatalf("'data' field is not an array")
	}
	if len(data) != 1 {
		t.Errorf("expected 'data' to contain 1 item, got %d", len(data))
	}
}

// mockTransport intercepts HTTP requests and redirects them to a mock server.
type mockTransport struct {
	base     http.RoundTripper
	mockHost string
}

// RoundTrip overrides the HTTP request's URL to redirect to the mock server.
func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Host = m.mockHost[len("http://"):]
	req.URL.Scheme = "http"
	return m.base.RoundTrip(req)
}
