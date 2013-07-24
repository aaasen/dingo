package dingo

import (
	"net/http"
)

var router *Router

func handle(w http.ResponseWriter, r *http.Request) {
	router.GetController(r).Respond(w, r)
}

func Run(routerPath string) {
	router = LoadRouter(routerPath)

	http.HandleFunc("/", handle)
	http.ListenAndServe(":"+Port, nil)
}
