package main

import "net/http"

func RootEndpoint() APIEndpoint {
	return APIEndpoint{
		Path: "/",
		Methods: []APIMethod{
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
