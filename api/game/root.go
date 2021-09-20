package game

import (
	"log"
	"net/http"

	models "github.com/nbelkaid/battlesnakes/api/models"
	"github.com/nbelkaid/battlesnakes/api/utils"
)

func info() models.BattlesnakeInfoResponse {
	log.Println("INFO")
	return models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "nbelkaid",
		Color:      "#888888",
		Head:       "default",
		Tail:       "default",
	}
}

//HandleIndex Respond with basic setting for Snake
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := info()

	err := utils.Respond(w, response)
	if err != nil {
		log.Printf("ERROR: Failed to encode info response, %s", err)
	}
}
