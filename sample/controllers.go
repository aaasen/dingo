package main

import (
	"github.com/aaasen/dingo"
	"net/http"
)

type IndexController struct{}

func (c IndexController) Respond(w http.ResponseWriter, r *http.Request, data map[string]string) {
	dingo.RenderTemplate(w, "index.html", nil)
}

type Page struct {
	Title string
}

type PageView struct{}

func (v PageView) Render(w http.ResponseWriter, data interface{}) {
	dingo.RenderTemplate(w, "page.html", data)
}

type PageController struct{}

func (c PageController) Respond(w http.ResponseWriter, r *http.Request, data map[string]string) {
	m := Page{"hello"}
	v := new(PageView)
	v.Render(w, m)
}
