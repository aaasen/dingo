package main

import (
	"github.com/aaasen/dingo"
)

func main() {
	devConf := dingo.Config{
		Port:        "8080",
		TemplateDir: "templates/",
		StaticDir:   "assets/",
		Routes:      routes,
	}

	server := dingo.New(devConf)

	server.Run()
}
