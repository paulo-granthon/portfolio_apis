package endpoints

import (
	"net/http"
	"server"
)

func UserEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/users",
			Methods: []server.Method{
				server.NewMethod("GET", GetUsers),
				server.NewMethod("POST", CreateUser),
			},
		},
		{
			Path: "/users/{id}",
			Methods: []server.Method{
				server.NewMethod("GET", GetUser),
				server.NewMethod("PUT", UpdateUser),
				server.NewMethod("DELETE", DeleteUser),
			},
		},
	}
}

func GetUsers(s server.Server, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func UpdateUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func DeleteUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	return nil
}
