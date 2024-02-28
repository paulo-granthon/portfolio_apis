package endpoints

import (
	"net/http"
	"server"
)

func CreateEndpoints() []server.Endpoint {
	return append(
		ProjectEndpoints(),
		RootEndpoint(),
	)
}

func RootEndpoint() server.Endpoint {
	return server.Endpoint{
		Path: "/",
		Methods: []server.Method{
			{
				Method: "GET",
				Func: func(s server.Server, w http.ResponseWriter, r *http.Request) error {
					w.Write([]byte("Hello, World!"))
					return nil
				},
			},
		},
	}
}
