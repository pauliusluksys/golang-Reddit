package handlers

import (
	"github.com/pauliusluksys/golang-Reddit/utils"
	"net/http"
)

// Greet request greet request
func Greet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	user := r.Context().Value("user").(string)
	// w.Write([]byte("hello, " + userID))
	utils.RespondWithJSON(w, http.StatusOK, "hello, "+user)
}
