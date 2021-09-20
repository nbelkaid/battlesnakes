package utils

import (
	"encoding/json"
	"net/http"
)

//Abs - Get absolute from int
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//Respond Add Header and Encore response
func Respond(w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}
