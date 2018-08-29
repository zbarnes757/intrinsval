package main

import (
	"intrinsval/pkg/server"
)

func main() {
	// var (
	// 	ticker               string
	// 	growthRate           float64
	// 	requiredRateOfReturn float64
	// 	riskFreeRate         float64
	// 	earnings             *iex.EarningsResponse
	// 	quote                *iex.QuoteResponse
	// )

	// getFlags(&ticker, &growthRate, &requiredRateOfReturn, &riskFreeRate)

	// fetchIEXData(ticker, earnings, quote)

	// // derived from http://www.moneychimp.com/articles/valuation/pe_ratio.htm with 9% discount rate
	// pe := float64(11.11)

	// data := &calculator.Data{
	// 	EPS:                  float64(earnings.Earnings[0].ActualEPS),
	// 	PE:                   pe,
	// 	GrowthRate:           growthRate,
	// 	RiskFreeRate:         riskFreeRate,
	// 	RequiredRateOfReturn: requiredRateOfReturn,
	// 	Quote:                quote,
	// }

	// calculator.RunFormulas(data)
	// fmt.Printf("The latest price is: $%.2f\n", quote.LatestPrice)

	s := server.New()

	s.RegisterAPIRoutes()

	s.Run()
}

// func getFlags(
// 	ticker *string,
// 	growthRate *float64,
// 	requiredRateOfReturn *float64,
// 	riskFreeRate *float64,
// ) {
// 	flag.StringVar(ticker, "ticker", "aapl", "a stock ticker symbol")

// 	flag.Float64Var(growthRate, "growthRate", 10.00, "Projected growth rate for next 5 years")

// 	flag.Float64Var(requiredRateOfReturn, "rateOfReturn", 5.00, "Your required rate of return")

// 	flag.Float64Var(
// 		riskFreeRate,
// 		"riskFreeRate",
// 		3.84,
// 		"risk-free rate (for us this comes from AAA corporate rate bond)",
// 	)

// 	flag.Parse()
// }

// func fetchIEXData(ticker string, earnings *iex.EarningsResponse, quote *iex.QuoteResponse) {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		e, err := iex.GetEarnings(ticker)
// 		if err != nil {
// 			panic(err)
// 		}

// 		earnings = e
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		q, err := iex.GetQuote(ticker)
// 		if err != nil {
// 			panic(err)
// 		}

// 		quote = q
// 	}()

// 	wg.Wait()
// }
