package server

import (
	"fmt"
	"net/http"
)

type Method struct {
	Method string
	Func   Func
}

func (e *Endpoint) getFunc(method string) (Func, error) {
	for _, m := range e.Methods {
		if m.Method != method {
			continue
		}
		return m.Func, nil
	}
	return nil, fmt.Errorf("Method %s not allowed", method)
}

type Func func(w http.ResponseWriter, r *http.Request) error

func NewHTTPHandlerFunc(apiEndpoint Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiFunc, err := apiEndpoint.getFunc(r.Method)

		if err != nil {
			WriteJSON(w, http.StatusMethodNotAllowed, Error{Error: err.Error()})
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
			return
		}

		if err := apiFunc(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, Error{Error: err.Error()})
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
