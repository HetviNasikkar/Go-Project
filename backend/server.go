package backend

import (
	"encoding/json"
	"net/http"
)

// Crypto prices data
func getCryptoPrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]map[string]string{
		"BTC":  {"price": "45000", "change": "+2.5%"},
		"ETH":  {"price": "3200", "change": "-1.2%"},
		"DOGE": {"price": "0.08", "change": "+5.0%"},
	}
	json.NewEncoder(w).Encode(data)
}

// Market trends data
func getMarketTrends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"trend":      "Bullish",
		"top_movers": "BTC, ETH, SOL",
	}
	json.NewEncoder(w).Encode(data)
}

// Crypto news data
func getCryptoNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := []map[string]string{
		{"title": "Bitcoin Hits New High", "source": "CoinDesk"},
		{"title": "Ethereum Upgrade Coming Soon", "source": "CryptoNews"},
	}
	json.NewEncoder(w).Encode(data)
}

// Setup routes
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/crypto-prices", getCryptoPrices)
	mux.HandleFunc("/api/market-trends", getMarketTrends)
	mux.HandleFunc("/api/news", getCryptoNews)
	return mux
}
