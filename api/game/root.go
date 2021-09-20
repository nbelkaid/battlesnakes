package game

import (
	"encoding/json"
	"log"
	"net/http"

	models "github.com/nbelkaid/battlesnakes/api/models"
)

func info() models.BattlesnakeInfoResponse {
	log.Println("INFO")
	return models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "",        // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Personalize
		Head:       "default", // TODO: Personalize
		Tail:       "default", // TODO: Personalize
	}
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := info()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode info response, %s", err)
	}
}
