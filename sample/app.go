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

	route := dingo.NewHandler("GET", "/", IndexController{})
	server.Router.AddHandler(route)

	pageRoute := dingo.NewHandler("GET", "/page", PageController{})
	server.Router.AddHandler(pageRoute)

	assetRoute := dingo.NewHandler("GET", "/assets/<path>/<name>.css", StaticController{"assets/"})
	server.Router.AddHandler(assetRoute)

	server.Run()
}
