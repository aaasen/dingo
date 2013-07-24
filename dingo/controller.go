package dingo

import (
	"fmt"
	"net/http"
)

type Controller interface {
	Respond(http.ResponseWriter, *http.Request)
}

type IndexController struct {
}

func (i IndexController) Respond(w http.ResponseWriter, r *http.Request) {
	fmt.Println("controller called")
}
