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

func end(state models.GameState) {
	tx := DB.BeginTx(context.Background(), nil)

	gameLog := models.GameLog{}

	//Get game to finish
	tx.Where(&models.GameLog{ID: state.Game.ID}).First(&gameLog)

	//Set data
	gameLog.End = time.Now()
	gameLog.NbTurn = state.Turn
	if len(state.Board.Snakes) == 1 && state.Board.Snakes[0].Head == state.You.Head {
		// I Won!
		gameLog.Position = 1
		gameLog.Won = true
	} else if len(state.Board.Snakes) == 1 {
		gameLog.Position = len(state.Board.Snakes) + 1
	}

	//Save & Commit
	if err := tx.Save(&gameLog).Error; err != nil {
		fmt.Println("Unable to Save gamelog")
		tx.Rollback()
	} else {
		tx.Commit()
	}

	log.Printf("%s END\n\n", state.Game.ID)
}

//Handler for /end endpoint - No response Needed - Store in Database some log about the game just ended
func HandleEnd(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}

	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode end json, %s", err)
		return
	}

	end(state)
}
