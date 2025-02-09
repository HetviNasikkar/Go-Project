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
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var data CryptoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	result := make(map[string]map[string]string)
	for key, value := range data {
		result[key] = map[string]string{
			"price":  fmt.Sprintf("%.2f", value.Price),
			"change": fmt.Sprintf("%.2f%%", value.Change),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/api/crypto-prices", getCryptoPrices)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
