package server

import (
	"net/http"

	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
)

// Server is the instance that holds all shared server logic
type Server struct {
	router *mux.Router
}

// New returns a new instance of a server
func New() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

// RegisterAPIRoutes will setup all the api routes for the front end to consume
func (server *Server) RegisterAPIRoutes() {
	api := server.router.PathPrefix("/api/v1").Subrouter()

	api.
		Path("/welcome").
		Methods("GET").
		HandlerFunc(server.jsonHandler)
}

// Run starts up the server on the desired port
func (server *Server) Run() {
	n := negroni.Classic()

	n.Use(negroni.NewStatic(http.Dir("web/dist")))

	n.UseHandler(server.router)

	n.Run()
}
