package dingo

import (
	"net/http"
)

type Server struct {
	Router *ARouter
}

// global configuration
var conf Config

func New(config Config) *Server {
	conf = config
	server := Server{Router: NewRouter()}
	return &server
}

func (server *Server) handle(w http.ResponseWriter, r *http.Request) {
	controller, err := server.Router.GetController(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		controller.Respond(w, r)
	}
}

func (server *Server) Run() {
	http.HandleFunc("/", server.handle)
	http.ListenAndServe(":"+conf.Port, nil)
}
