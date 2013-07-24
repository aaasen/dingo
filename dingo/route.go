package dingo

import (
	"net/http"
)

type Route struct {
	method string
	path   string
	action string
}

type Router struct {
	routes []*Route
}

func (router *Router) GetController(r *http.Request) Controller {
	cont := IndexController{}
	return cont
}

func LoadRouter(conf string) *Router {
	router := new(Router)
	routes := make([]*Route, 0)

	route := &Route{"GET", "/", "IndexController"}
	routes = append(routes, route)

	router.routes = routes

	return router
}
