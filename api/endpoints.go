package main

import (
	"net/http"
	"server"
)

func CreateEndpoints() []server.Endpoint {
	return append(
		[]server.Endpoint{},
		RootEndpoint(),
	)
}

func RootEndpoint() server.Endpoint {
	return server.Endpoint{
		Path: "/",
		Methods: []server.Method{
			{
				Method: "GET",
				Func: func(w http.ResponseWriter, r *http.Request) error {
					w.Write([]byte("Hello, World!"))
					return nil
				},
			},
		},
	}
}
