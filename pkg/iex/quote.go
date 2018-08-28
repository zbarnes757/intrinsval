package iex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// QuoteResponse is the parsed json body of an quote request
type QuoteResponse struct {
	Symbol      string  `json:"symbol"`
	CompanyName string  `json:"companyName"`
	LatestPrice float32 `json:"latestPrice"`
	PeRatio     float32 `json:"peRatio"`
}

// GetQuote will return the relevant stock information for a given ticker
func GetQuote(ticker string) (*QuoteResponse, error) {
	requestURL := fmt.Sprintf("%s/stock/%s/quote", iexURI, ticker)
	resp, err := http.Get(requestURL)
	if err != nil {
		return &QuoteResponse{}, err
	}

	var quote QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		return &QuoteResponse{}, err
	}

	return &quote, nil

}
