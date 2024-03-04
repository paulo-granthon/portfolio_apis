package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRequestParam(r *http.Request, key string) (*string, error) {
	param := mux.Vars(r)[key]

	if key == "" {
		return nil, fmt.Errorf("Parameter %s not found", key)
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
	log.Println(err)
	return WriteJSON(w, code, Error{Error: msg})
}
