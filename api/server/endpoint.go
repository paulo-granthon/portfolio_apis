package server

import (
	"net/http"
)

type Endpoint struct {
	Path    string
	Methods []Method
}

func NewEndpoint(path string, methods []Method) *Endpoint {
	return &Endpoint{
		Path:    path,
		Methods: methods,
	}
}

func (endpoint Endpoint) Create(server Server) http.HandlerFunc {
	return NewHTTPHandlerFunc(server, endpoint)
}
