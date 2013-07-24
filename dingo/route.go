package dingo

import (
	"net/http"
)

type Route struct {
	Method     string
	Path       string
	Controller Controller
}

type Router struct {
	Routes []*Route
}

func (router *Router) GetController(r *http.Request) Controller {
	cont := IndexController{}
	return cont
}

func NewRouter() *Router {
	router := new(Router)
	router.Routes = make([]*Route, 0)
	return router
}

func (router *Router) AddRoute(route *Route) {
	router.Routes = append(router.Routes, route)
}

// func processRouterConfig(configPath string) ([]byte, error) {
// 	content, err := ioutil.ReadFile(configPath)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return content, nil
// }

// var commentExtractor = regexp.MustCompile("(?m)^([^#\n]+)$")
// var whitespaceExtractor = regexp.MustCompile("[ \t\r\n\v\f]+")
// var methodValidator = regexp.MustCompile("[A-Za-z]+")
// var pathValidator = regexp.MustCompile("[A-Za-z0-9/<>\\.]+")
// var controllerValidator = regexp.MustCompile("[A-Za-z0-9\\.]+")

// func LoadRouter(path string) *Router {
// 	conf, _ := processRouterConfig(path)

// 	router := new(Router)
// 	routes := make([]*Route, 0)

// 	for _, value := range commentExtractor.FindAllString(string(conf), -1) {
// 		r := whitespaceExtractor.Split(value, -1)

// 		if len(r) == 3 {
// 			if methodValidator.MatchString(r[0]) &&
// 				pathValidator.MatchString(r[1]) &&
// 				controllerValidator.MatchString(r[2]) {

// 				route := &Route{strings.ToUpper(r[0]), r[1], r[2]}
// 				routes = append(routes, route)
// 				fmt.Println(route)
// 			}
// 		}
// 	}

// 	router.routes = routes

// 	return router
// }
