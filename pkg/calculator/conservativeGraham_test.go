package calculator

import (
	"intrinsval/pkg/iex"
	"testing"
)

func Test_conservativeGrahamsIntrinsicValue(t *testing.T) {
	type args struct {
		data   *Data
		result float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Given a set of data",
			args{
				data: &Data{
					GrowthRate:           float64(10.00),
					PE:                   float64(10.00),
					EPS:                  float64(10.00),
					RequiredRateOfReturn: float64(10.00),
					RiskFreeRate:         float64(10.00),
					Quote: &iex.QuoteResponse{
						Symbol:      "aapl",
						CompanyName: "my company",
						LatestPrice: float32(5.00),
						PeRatio:     float32(5.00),
					},
				},
				result: float64(190.00),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := conservativeGrahamsIntrinsicValue(tt.args.data)
			if res != tt.args.result {
				t.Errorf("Result was incorrect, got: %.2f, want: %.2f.", res, tt.args.result)
			}
		})
	}
}
