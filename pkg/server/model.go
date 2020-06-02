package server

import "net/http"

type Endpoint struct {
	Method  string
	Path    string
	Handler func(*http.Request) http.Response
}
