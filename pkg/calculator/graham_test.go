package calculator

import (
	"intrinsval/pkg/iex"
	"testing"
)

func Test_grahamsIntrinsicValue(t *testing.T) {
	type args struct {
		data *Data
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Given a set of data",
			args: args{
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
			},
			want: float64(300.00),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grahamsIntrinsicValue(tt.args.data); got != tt.want {
				t.Errorf("grahamsIntrinsicValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
