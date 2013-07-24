package dingo

import (
	"net/http"
	"regexp"
)

type Route struct {
	method     string
	path       string
	pathRegex  *regexp.Regexp
	controller Controller
}

type Router struct {
	routes []*Route
}

func NewRoute(method string, path string, controller Controller) *Route {
	route := &Route{
		method:     method,
		path:       path,
		pathRegex:  regexp.MustCompile("^" + path + "/?$"),
		controller: controller,
	}

	return route
}

func NewRouter() *Router {
	router := new(Router)
	router.routes = make([]*Route, 0)
	return router
}

func (router *Router) GetController(r *http.Request) Controller {
	for _, value := range router.routes {
		//check that method is correct (GET, POST)
		if value.pathRegex.MatchString(r.URL.Path) {
			return value.controller
		}
	}

	return new(Controller404)
}

func (router *Router) AddRoute(route *Route) {
	router.routes = append(router.routes, route)
}
