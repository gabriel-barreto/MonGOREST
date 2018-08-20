package functions

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseWithError is used to send a JSON response content
func ResponseWithError(w http.ResponseWriter, code int, message string) {
	payload := struct {
		Message string `json:"message"`
	}{message}
	ResponseWithJSON(w, code, payload)
}

// ResponseWithJSON is used to send a JSON response content
func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
