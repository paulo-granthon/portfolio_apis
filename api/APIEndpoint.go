package main

import "net/http"

type APIEndpoint struct {
	Path    string
	Methods []APIMethod
}

func NewAPIEndpoint(path string, methods []APIMethod) *APIEndpoint {
	return &APIEndpoint{
		Path:    path,
		Methods: methods,
	}
}

func (e APIEndpoint) Create() http.HandlerFunc {
	return NewHTTPHandlerFunc(e)
}
