package handler

import (
	"fmt"
	"net/http"
)

type IHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

func respondJSON(w http.ResponseWriter, code int, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, data)
}
