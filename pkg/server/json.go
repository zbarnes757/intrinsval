package server

import (
	"encoding/json"
	"net/http"
)

type hello struct {
	Welcome string `json:"welcome"`
}

func (server *Server) jsonHandler(w http.ResponseWriter, r *http.Request) {
	resp := &hello{"zac"}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
