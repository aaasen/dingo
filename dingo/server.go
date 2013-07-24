package dingo

import (
	"net/http"
)

type Server struct {
	Router *Router
}

// global configuration
var conf Config

func New(config Config) *Server {
	conf = config
	server := Server{Router: NewRouter()}
	return &server
}

func (server *Server) handle(w http.ResponseWriter, r *http.Request) {
	server.Router.GetController(r).Respond(w, r)
}

func (server *Server) Run() {
	http.HandleFunc("/", server.handle)
	http.ListenAndServe(":"+conf.Port, nil)
}
