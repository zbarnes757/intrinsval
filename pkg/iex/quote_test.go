package iex

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPI_GetQuote(t *testing.T) {
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
		want    *QuoteResponse
		wantErr bool
	}{
		{
			"successful call",
			fields{serviceEndpoint{}},
			args{"aapl"},
			&QuoteResponse{
				"aapl",
				"apple compant",
				float32(50.00),
				float32(2.5),
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

			got, err := api.GetQuote(tt.args.ticker)

			if (err != nil) != tt.wantErr {
				t.Errorf("API.GetQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("API.GetQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}
