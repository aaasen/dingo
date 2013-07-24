package main

import (
	"github.com/aaasen/dingo"
)

func main() {
	devConf := dingo.Config{
		Port:        "8080",
		TemplateDir: "templates/",
	}

	server := dingo.New(devConf)

	route := dingo.NewRoute("GET", "/", IndexController{})
	server.Router.AddRoute(route)

	pageRoute := dingo.NewRoute("GET", "/page", PageController{})
	server.Router.AddRoute(pageRoute)

	server.Run()
}
