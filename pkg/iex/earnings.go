package iex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// EarningsResponse is the full json body of an earnings request
type EarningsResponse struct {
	Symbol   string    `json:"symbol"`
	Earnings []Earning `json:"earnings"`
}

// Earning represents one quarter's earnings
type Earning struct {
	ActualEPS     float32 `json:"actualEPS"`
	ConsensusEPS  float32 `json:"consensusEPS"`
	EstimatedEPS  float32 `json:"estimatedEPS"`
	EPSReportDate string  `json:"EPSReportDate"`
	FiscalPeriod  string  `json:"fiscalPeriod"`
	FiscalEndDate string  `json:"fiscalEndDate"`
}

// GetEarnings pulls data from the four most recent reported quarters
func (api *API) GetEarnings(ticker string) (*EarningsResponse, error) {
	requestURL := fmt.Sprintf("%s/stock/%s/earnings", api.url, ticker)
	resp, err := http.Get(requestURL)
	if err != nil {
		return &EarningsResponse{}, err
	}

	var earnings EarningsResponse
	err = json.NewDecoder(resp.Body).Decode(&earnings)
	if err != nil {
		return &EarningsResponse{}, err
	}

	return &earnings, nil
}
