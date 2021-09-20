package game

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sort"

	models "github.com/nbelkaid/battlesnakes/api/models"
	"github.com/nbelkaid/battlesnakes/api/utils"
)

//Check if coord are next to one Food
func isFoodNextTo(coord models.Coord, foods []models.Coord) bool {
	for _, food := range foods {
		if coord.IsNextTo(food) {
			return true
		}
	}
	return false
}

//Base on Head coordinates and a obstacle Coordinate
func removeNotPossibleMoveFromObstacle(possibleMoves map[string]bool, myHead, obstacle models.Coord) map[string]bool {
	if obstacle.X < myHead.X {
		possibleMoves["left"] = false
	} else if obstacle.X > myHead.X {
		possibleMoves["right"] = false
	} else if obstacle.Y < myHead.Y {
		possibleMoves["down"] = false
	} else if obstacle.Y > myHead.Y {
		possibleMoves["up"] = false
	}

	return possibleMoves
}

//Based on Head Coordinates we remove not possible move to avoid Wall Collision
func removeNotPossibleMoveFromWalls(possibleMoves map[string]bool, myHead models.Coord, width, height int) map[string]bool {
	//Remove 1 because index start at 0
	width = width - 1
	height = height - 1

	if myHead.X == 0 {
		possibleMoves["left"] = false
	}
	if myHead.X == width {
		possibleMoves["right"] = false
	}
	if myHead.Y == 0 {
		possibleMoves["down"] = false
	}
	if myHead.Y == height {
		possibleMoves["up"] = false
	}

	return possibleMoves
}

//Will respond with possible move with wall collision constraint and self and other snakes collision constraint
func getPossibleMoves(state models.GameState) map[string]bool {
	myHead := state.You.Head

	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	//Apply walls collision constraint
	possibleMoves = removeNotPossibleMoveFromWalls(possibleMoves, myHead, state.Board.Width, state.Board.Height)

	//Apply myself and Other snakes collission Constraint`
	for _, snake := range state.Board.Snakes {
		for _, coord := range snake.Body {
			//ToImprove : No need to considerate tail untill no food is next to head of current snake
			if myHead.IsNextTo(coord) {
				possibleMoves = removeNotPossibleMoveFromObstacle(possibleMoves, myHead, coord)
			}

		}
	}

	return possibleMoves
}

func moveToGetCloserToDest(safeMoves []string, head, dest models.Coord) string {
	for _, move := range safeMoves {
		switch move {
		case "left":
			if dest.X < head.X {
				return move
			}
		case "right":
			if dest.X > head.X {
				return move
			}
		case "up":
			if dest.Y > head.Y {
				return move
			}
		case "down":
			if dest.Y < head.Y {
				return move
			}
		}
	}

	return ""
}

// This function is called on every turn of a game.
// valid moves are "up", "down", "left", or "right".
func move(state models.GameState) models.BattlesnakeMoveResponse {
	myHead := state.You.Head
	possibleMoves := getPossibleMoves(state)

	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
		return models.BattlesnakeMoveResponse{
			Move: nextMove,
		}
	}

	//Here we want to order food by distance from our Head ASC
	foods := state.Board.Food
	sort.SliceStable(foods, func(i, j int) bool {
		dXi := utils.Abs(myHead.X - foods[i].X)
		dYi := utils.Abs(myHead.Y - foods[i].Y)
		totalI := dXi + dYi

		dXj := utils.Abs(myHead.X - foods[j].X)
		dYj := utils.Abs(myHead.Y - foods[j].Y)
		totalJ := dXj + dYj

		return totalI < totalJ
	})

	for _, food := range foods {
		nextMove = moveToGetCloserToDest(safeMoves, myHead, food)
		if nextMove != "" {
			log.Printf("%s MOVE %d: Food On The Way! Moving %s\n", state.Game.ID, state.Turn, nextMove)
			return models.BattlesnakeMoveResponse{
				Move: nextMove,
			}
		}
	}

	nextMove = safeMoves[rand.Intn(len(safeMoves))]
	log.Printf("%s MOVE %d: Random Move! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	return models.BattlesnakeMoveResponse{
		Move: nextMove,
	}
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode move json, %s", err)
		return
	}

	response := move(state)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode move response, %s", err)
		return
	}
}
