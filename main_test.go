package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleLTPRequestIntegration(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handleLTPRequest))
	defer ts.Close()

	oldPort := os.Getenv("PORT")
	defer os.Setenv("PORT", oldPort)
	os.Setenv("PORT", "8081")

	resp, err := http.Get(ts.URL + "/api/v1/ltp")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var prices []LTPResponse
	if err := json.NewDecoder(resp.Body).Decode(&prices); err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	expectedNumPrices := 3
	if len(prices) != expectedNumPrices {
		t.Errorf("Expected %d prices, got %d", expectedNumPrices, len(prices))
	}

	for _, price := range prices {
		if price.Pair == "" {
			t.Errorf("Empty pair found in response")
		}
		if price.Price == "" {
			t.Errorf("Empty price found in response")
		}
	}
}
