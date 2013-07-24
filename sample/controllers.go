package main

import (
	"fmt"
	"net/http"
)

type IndexController struct{}

func (c IndexController) Respond(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index!")
}
