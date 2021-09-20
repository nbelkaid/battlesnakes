package game

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	models "github.com/nbelkaid/battlesnakes/api/models"
)

func start(state models.GameState) {

	tx := DB.BeginTx(context.Background(), nil)

	gameLog := models.GameLog{}

	//Set Data
	gameLog.ID = state.Game.ID
	gameLog.Start = time.Now()
	gameLog.NbSnake = len(state.Board.Snakes)
	gameLog.NbTurn = 0
	gameLog.Ruleset = state.Game.Ruleset.Name
	gameLog.Width = state.Board.Width
	gameLog.Height = state.Board.Height
	gameLog.AlgVersion = 1

	//Create Row
	if err := tx.Select("ID", "Start", "NbSnake", "NbTurn", "Ruleset", "Width", "Height", "AlgVersion").Create(&gameLog).Error; err != nil {
		fmt.Println("Unable to create gamelog")
		tx.Rollback()
	} else {
		tx.Commit()
	}

	log.Printf("START\n")
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode start json, %s", err)
		return
	}

	start(state)

	// Nothing to respond with here
}
