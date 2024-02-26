package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRequestParam(r *http.Request, key string) (any, error) {
	param := mux.Vars(r)[key]

	if key == "" {
		return nil, fmt.Errorf("Parameter %s not found", key)
	}

	return param, nil
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

type Error struct {
	Error string
}

func SendError(w http.ResponseWriter, error error) error {
	return WriteJSON(w, http.StatusBadRequest, Error{Error: error.Error()})
}