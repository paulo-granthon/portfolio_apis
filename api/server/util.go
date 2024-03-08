package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ztrue/tracerr"
)

func GetRequestParam(r *http.Request, key string) (*string, error) {
	param := mux.Vars(r)[key]

	if key == "" {
		return nil, tracerr.Errorf("Parameter %s not found", key)
	}

	return &param, nil
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func ReadJSON[T comparable](r *http.Request, value *T) error {
	return json.NewDecoder(r.Body).Decode(value)
}

type Error struct {
	Error string
}

func SendError(
	w http.ResponseWriter,
	err error,
	code int,
	msg string,
) error {
	tracerr.PrintSourceColor(err, 1)
	return WriteJSON(w, code, Error{Error: msg})
}
