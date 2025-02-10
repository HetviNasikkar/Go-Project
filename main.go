package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CryptoResponse represents the API response structure
type CryptoResponse map[string]struct {
	Price  float64 `json:"usd"`
	Change float64 `json:"usd_24h_change"`
}

// Handler function to fetch crypto prices
func getCryptoPrices(w http.ResponseWriter, r *http.Request) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,cardano&vs_currencies=usd&include_24hr_change=true"

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch data: %v", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("API returned non-200 status: %d", resp.StatusCode)
		http.Error(w, "API returned non-200 status", http.StatusInternalServerError)
		return
	}

	var data CryptoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding response: %v", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	result := make(map[string]map[string]float64)
	for key, value := range data {
		result[key] = map[string]float64{
			"price":  value.Price,
			"change": value.Change,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Enable CORS
	json.NewEncoder(w).Encode(result)
}

// News handler
// News handler
func getCryptoNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Enable CORS
	data := []map[string]string{
		{
			"title":  "Bitcoin Hits New High",
			"source": "CoinDesk",
			"url":    "https://www.coindesk.com/bitcoin-hits-new-high",
		},
		{
			"title":  "Ethereum Upgrade Coming Soon",
			"source": "CryptoNews",
			"url":    "https://www.cryptonews.com/ethereum-upgrade-coming-soon",
		},
	}
	json.NewEncoder(w).Encode(data)
}

// Market Trends handler
// Market Trends handler
func getMarketTrends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Enable CORS
	data := []map[string]string{
		{
			"trend":      "Bullish",
			"top_movers": "BTC, ETH, SOL",
			"description": "The market is showing strong upward momentum, with Bitcoin leading the charge.",
			"icon":        "📈", // Emoji for visual representation
		},
		{
			"trend":      "Bearish",
			"top_movers": "DOGE, SHIB, XRP",
			"description": "The market is experiencing a downturn, with meme coins underperforming.",
			"icon":        "📉", // Emoji for visual representation
		},
		{
			"trend":      "Neutral",
			"top_movers": "ADA, DOT, AVAX",
			"description": "The market is stable, with no significant upward or downward movement.",
			"icon":        "➖", // Emoji for visual representation
		},
		{
			"trend":      "Volatile",
			"top_movers": "LUNA, FTT, MATIC",
			"description": "The market is highly volatile, with rapid price fluctuations.",
			"icon":        "🌪️", // Emoji for visual representation
		},
	}
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register backend routes
	mux.HandleFunc("/api/crypto", getCryptoPrices)
	mux.HandleFunc("/api/market-trends", getMarketTrends)
	mux.HandleFunc("/api/news", getCryptoNews)

	// Start server
	port := ":8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}