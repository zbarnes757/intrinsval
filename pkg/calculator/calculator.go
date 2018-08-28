package calculator

import "intrinsval/pkg/iex"

// Data is information required for calculations
type Data struct {
	GrowthRate           float64
	EPS                  float64
	PE                   float64
	RequiredRateOfReturn float64
	RiskFreeRate         float64
	Quote                *iex.QuoteResponse
}

// RunFormulas will calculate all the formulas
func RunFormulas(data *Data) {
	grahamsIntrinsicValue(data)
	conservativeGrahamsIntrinsicValue(data)
}
