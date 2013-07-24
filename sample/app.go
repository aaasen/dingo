package main

import (
	"github.com/aaasen/dingo"
)

func main() {
	server := dingo.New()

	route := &dingo.Route{
		Method:     "GET",
		Path:       "/",
		Controller: dingo.IndexController{},
	}

	server.Router.AddRoute(route)

	server.Run()
}
