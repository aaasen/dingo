package dingo

import (
	"fmt"
	"net/http"
)

type Controller interface {
	Respond(http.ResponseWriter, *http.Request)
}

type Controller404 struct{}

func (c Controller404) Respond(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "404 error: file not found")
}
