package main

import (
	"github.com/aaasen/dingo"
)

func main() {
	server := dingo.New()

	route := dingo.NewRoute(
		"GET",
		"/",
		IndexController{},
	)

	server.Router.AddRoute(route)

	server.Run()
}
