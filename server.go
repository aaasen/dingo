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
	server.Router.SetHandlers(config.Routes)
	return &server
}

func (server *Server) handle(w http.ResponseWriter, r *http.Request) {
	server.Router.Route(w, r)
}

func (server *Server) Run() {
	http.Handle("/"+conf.StaticDir,
		http.StripPrefix("/"+conf.StaticDir,
			http.FileServer(http.Dir(conf.StaticDir))))

	http.HandleFunc("/", server.handle)
	http.ListenAndServe(":"+conf.Port, nil)
}
