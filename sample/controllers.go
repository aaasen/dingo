package main

import (
	"github.com/aaasen/dingo"
	"io/ioutil"
	"net/http"
)

type IndexController struct{}

func (c IndexController) Respond(w http.ResponseWriter, r *http.Request) {
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

func (c PageController) Respond(w http.ResponseWriter, r *http.Request) {
	m := Page{"hello"}
	v := new(PageView)
	v.Render(w, m)
}

type StaticController struct {
	rootPath string
}

func (c StaticController) Respond(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile(r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Write(content)
}
