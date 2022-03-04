package v1

import (
	"github.com/pauliusluksys/golang-Reddit/utils"
	"net/http"
)

// Greet request greet request
func Greet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value("user_id").(string)
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("hello, " + userID))
	utils.RespondWithJSON(w, "hello, "+userID)
}
