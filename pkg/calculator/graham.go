package calculator

import (
	"fmt"
)

// GrahamsIntrinsicValue calculates based on Ben Graham's formula
func grahamsIntrinsicValue(data *Data) {
	/*
		V = (EPS * (PE + 2G) * RROR) / RFR
		V is intrinsic value
		EPS is earnings per share
		PE is PE ratio of a stock that has 0% growth
		G is growth rate for the next 7 to 10 years
		RROR is required rate of return
		RFR is risk-free rate (for us this comes from AAA corporate rate bond)
	*/

	val := (data.EPS * (data.PE + 2*data.GrowthRate) * data.RequiredRateOfReturn) / data.RiskFreeRate
	fmt.Printf(
		"Based on Ben Graham's formula, the intrinsic value of %s (%s) is: $%.2f\n",
		data.Quote.CompanyName,
		data.Quote.Symbol,
		val)
}
