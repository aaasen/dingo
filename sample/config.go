package main

import (
	"github.com/aaasen/dingo"
)

var config = dingo.Config{
	Port:        "8080",
	TemplateDir: "templates/",
	StaticDir:   "assets/",
	Routes:      routes,
}
