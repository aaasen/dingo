package main

import (
	"github.com/aaasen/dingo"
)

func main() {
	devConf := dingo.Config{
		Port:        "8080",
		TemplateDir: "templates/",
		StaticDir:   "assets/",
	}

	server := dingo.New(devConf)

	route := dingo.NewHandler("GET", "/", IndexController{})
	server.Router.AddHandler(route)

	pageRoute := dingo.NewHandler("GET", "/page", PageController{})
	server.Router.AddHandler(pageRoute)

	server.Run()
}
