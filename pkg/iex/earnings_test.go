package iex

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_api_GetEarnings(t *testing.T) {
	type fields struct {
		serviceEndpoint serviceEndpoint
	}

	type args struct {
		ticker string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *EarningsResponse
		wantErr bool
	}{
		{
			"successful call",
			fields{serviceEndpoint{}},
			args{"aapl"},
			&EarningsResponse{
				"aapl",
				[]Earning{
					Earning{
						float32(3.00),
						float32(3.00),
						float32(3.00),
						"2018-12-27",
						"Q4",
						"Q4",
					},
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.serviceEndpoint.url = httptest.
				NewServer(
					http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(tt.want)
					})).URL

			api := &API{
				serviceEndpoint: tt.fields.serviceEndpoint,
			}

			got, err := api.GetEarnings(tt.args.ticker)

			if (err != nil) != tt.wantErr {
				t.Errorf("api.GetEarnings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("api.GetEarnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
