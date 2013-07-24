package main

import (
	"github.com/aaasen/dingo"
)

func main() {
	devConf := dingo.Config{
		Port: "8080",
	}

	server := dingo.New(devConf)

	route := dingo.NewRoute(
		"GET",
		"/",
		IndexController{},
	)

	server.Router.AddRoute(route)

	server.Run()
}
