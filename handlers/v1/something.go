package v1

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetSomething(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal("something is going on here")
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}
}
