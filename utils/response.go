package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, v interface{}) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		fmt.Println(err.Error())
	}
}
