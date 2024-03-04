package server

import (
	"fmt"
	"net/http"
)

type Method struct {
	Method string
	Func   Func
}

func NewMethod(method string, f Func) Method {
	return Method{
		Method: method,
		Func:   f,
	}
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

type Func func(s Server, w http.ResponseWriter, r *http.Request) error

func NewHTTPHandlerFunc(s Server, apiEndpoint Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiFunc, err := apiEndpoint.getFunc(r.Method)

		if err != nil {
			SendError(
				w, err, http.StatusMethodNotAllowed,
				fmt.Sprintf("method %s not allowed", r.Method),
			)
			return
		}

		if err := apiFunc(s, w, r); err != nil {
			SendError(
				w, err, http.StatusInternalServerError,
				fmt.Sprintf("error processing request: %s", err.Error()),
			)
			return
		}
	}
}
