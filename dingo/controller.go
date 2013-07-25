package dingo

import (
	"net/http"
)

type Controller interface {
	Respond(http.ResponseWriter, *http.Request, map[string]string)
}
