package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestServer_jsonHandler(t *testing.T) {
	type fields struct {
		router *mux.Router
	}

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"valid request",
			fields{},
			args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/welcome", nil)
			w := httptest.NewRecorder()

			new(Server).jsonHandler(w, req)
			if resp := w.Result(); resp.StatusCode != 200 {
				t.Errorf("jsonHandler status code = %v, wanted %v", resp.StatusCode, 200)
			}
		})
	}
}
