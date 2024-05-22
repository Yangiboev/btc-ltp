package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type LTPResponse struct {
	Pair  string `json:"pair"`
	Price string `json:"price"`
}

type KrakenTickerResponse struct {
	Error  []string              `json:"error"`
	Result map[string]TickerInfo `json:"result"`
}

type TickerInfo struct {
	Close []string `json:"c"`
}

func main() {
	http.HandleFunc("/api/v1/ltp", handleLTPRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleLTPRequest(w http.ResponseWriter, r *http.Request) {
	pairs := []string{
		"BTC/USD",
		"BTC/CHF",
		"BTC/EUR",
	}
	prices := make([]LTPResponse, 0, len(pairs))

	var wg sync.WaitGroup
	mu := &sync.Mutex{}
	for _, pair := range pairs {
		wg.Add(1)
		go func(pair string) {
			defer wg.Done()
			price, err := fetchLTP(pair)
			if err != nil {
				log.Printf("Failed to fetch LTP for %s: %v", pair, err)
				return
			}
			mu.Lock()
			prices = append(prices, LTPResponse{Pair: pair, Price: price})
			mu.Unlock()
		}(pair)
	}

	wg.Wait()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prices); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fetchLTP(pair string) (string, error) {
	apiURL := fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=%s", pair)
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("kraken API request failed with status %s", resp.Status)
	}

	var data KrakenTickerResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding response for %s: %v", pair, err)
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	if len(data.Error) > 0 {
		log.Printf("Kraken API error for %s: %v", pair, data.Error)
		return "", fmt.Errorf("kraken API error: %v", data.Error)
	}

	tickerInfo, ok := data.Result[pair]
	if !ok || len(tickerInfo.Close) < 1 {
		log.Printf("Invalid response for %s", pair)
		return "", fmt.Errorf("invalid response")
	}

	return tickerInfo.Close[0], nil
}
