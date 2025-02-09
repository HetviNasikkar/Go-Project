package handlers

import (
    "encoding/json"
    "net/http"
    "cryptoapp/backend/fetch" // Corrected import path
)

// GetCryptoPrices handles the API request for fetching cryptocurrency prices.
func GetCryptoPrices(w http.ResponseWriter, r *http.Request) {
    data, err := fetch.FetchCryptoPrices()
    if err != nil {
        http.Error(w, "Failed to fetch prices", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}