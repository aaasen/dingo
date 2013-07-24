package dingo

import (
	"net/http"
)

type Server struct {
	Router *Router
	config Config
}

func New(config Config) *Server {
	server := Server{Router: NewRouter(), config: config}
	return &server
}

func (server *Server) handle(w http.ResponseWriter, r *http.Request) {
	server.Router.GetController(r).Respond(w, r)
}

func (server *Server) Run() {
	http.HandleFunc("/", server.handle)
	http.ListenAndServe(":"+server.config.Port, nil)
}
