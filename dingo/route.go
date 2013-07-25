package dingo

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Handler interface {
	// returns whether or not the Handler wants to act on the give request
	// can take into account path requested, method (GET, POST), etc.
	SatisfiesRequest(*http.Request) bool
}

type Router interface {
	// routes the given request to a controller
	GetController(*http.Request) Controller
}

type AHandler struct {
	method     string
	path       string
	pathRegex  *regexp.Regexp
	controller Controller
}

type ARouter struct {
	routes []*AHandler
}

func NewRouter() *ARouter {
	router := new(ARouter)
	router.routes = make([]*AHandler, 0)
	return router
}

type ControllerNotFoundError struct {
	request *http.Request
}

func (e ControllerNotFoundError) Error() string {
	return "no controller found for path: " + e.request.URL.Path
}

func (router *ARouter) GetController(r *http.Request) (Controller, error) {
	for _, handler := range router.routes {
		if handler.SatisfiesRequest(r) {
			return handler.controller, nil
		}
	}

	return nil, ControllerNotFoundError{r}
}

func (router *ARouter) AddHandler(handler *AHandler) {
	router.routes = append(router.routes, handler)
}

func NewHandler(method string, path string, controller Controller) *AHandler {
	varReplacer := regexp.MustCompile("<[A-Za-z0-9\\.-]+>")
	pathRegex := varReplacer.ReplaceAllString(path, "(.+?)")
	pathRegex = "^" + pathRegex + "/?$"

	route := &AHandler{
		method:     method,
		path:       path,
		pathRegex:  regexp.MustCompile(pathRegex),
		controller: controller,
	}

	return route
}

func (handler *AHandler) SatisfiesRequest(r *http.Request) bool {
	return handler.pathRegex.MatchString(r.URL.Path) &&
		strings.EqualFold(r.Method, handler.method)
}
