package controllers

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, code int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
