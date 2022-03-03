package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(res http.ResponseWriter, code int, message string) {
	respondWithJSON(res, code, map[string]string{"error": message})
}

func respondWithJSON(res http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write(response)
}
