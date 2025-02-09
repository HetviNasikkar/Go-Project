package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

// CryptoPrice represents the structure of the API response
type CryptoPrice struct {
	Price string `json:"price"`
}

// FetchCryptoPrices fetches real-time crypto prices from an API
func FetchCryptoPrices() ([]byte, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch crypto prices: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Convert JSON response into struct (optional)
	var priceData map[string]map[string]float64
	if err := json.Unmarshal(body, &priceData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	// Reformat response as needed
	responseJSON, err := json.Marshal(priceData)
	if err != nil {
		return nil, fmt.Errorf("failed to encode JSON: %v", err)
	}

	return responseJSON, nil
}
