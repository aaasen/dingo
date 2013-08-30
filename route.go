package dingo

import (
	"net/http"
	"regexp"
	"strings"
)

type Handler interface {
	// returns whether or not the Handler wants to act on the give request
	// can take into account path requested, method (GET, POST), etc.
	SatisfiesRequest(*http.Request) bool

	// handles a given request, generally by passing it to a controller
	Handle(*http.Request)
}

type Router interface {
	// routes the given request to a controller
	GetController(*http.Request) Controller

	// routes a request to its appropriate handler
	Route(*http.Request)
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

func (router *ARouter) Route(w http.ResponseWriter, r *http.Request) {
	for _, handler := range router.routes {
		if handler.SatisfiesRequest(r) {
			handler.HandleRequest(w, r)
			return
		}
	}

	http.Error(w, "file not found", http.StatusNotFound)
}

func (router *ARouter) AddHandler(handler *AHandler) {
	router.routes = append(router.routes, handler)
}

func (router *ARouter) SetHandlers(handlers []*AHandler) {
	router.routes = handlers
}

var varReplacer = regexp.MustCompile("<([A-Za-z0-9\\.-]+?)>")

func NewHandler(method string, path string, controller Controller) *AHandler {
	pathRegex := regexp.MustCompile("^" +
		varReplacer.ReplaceAllString(path, "(.+)") + "/?$")

	route := &AHandler{
		method:     method,
		path:       path,
		pathRegex:  pathRegex,
		controller: controller,
	}

	return route
}

func (handler *AHandler) SatisfiesRequest(r *http.Request) bool {
	return handler.pathRegex.MatchString(r.URL.Path) &&
		strings.EqualFold(r.Method, handler.method)
}

func (handler *AHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	names := varReplacer.FindAllStringSubmatch(handler.path, -1)
	values := handler.pathRegex.FindAllStringSubmatch(r.URL.Path, -1)[0][1:]

	if len(names) != len(values) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	m := make(map[string]string)

	for i := 0; i < len(names); i++ {
		m[names[i][1]] = values[i]
	}

	handler.controller.Respond(w, r, m)
}
