package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

type APIError struct {
	Error string
}

type APIMethod struct {
	Method string
	Func   APIFunc
}

func (e *APIEndpoint) getAPIFunc(method string) (APIFunc, error) {
	for _, m := range e.Methods {
		if m.Method != method {
			continue
		}
		return m.Func, nil
	}
	return nil, fmt.Errorf("Method %s not allowed", method)
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func NewHTTPHandlerFunc(apiEndpoint APIEndpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiFunc, err := apiEndpoint.getAPIFunc(r.Method)

		if err != nil {
			WriteJSON(w, http.StatusMethodNotAllowed, APIError{Error: err.Error()})
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
			return
		}

		if err := apiFunc(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
