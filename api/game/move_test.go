package game

import (
	"testing"

	models "github.com/nbelkaid/battlesnakes/api/models"
	"github.com/stretchr/testify/assert"
)

func TestIsFoodNextTo(t *testing.T) {
	foods := []models.Coord{
		{X: 42, Y: 42},
		{X: 10, Y: 10},
	}

	foodsEmpty := []models.Coord{}

	tt := []struct {
		Name           string
		head           models.Coord
		food           []models.Coord
		expectedResult bool
	}{
		{"true with x + 1", models.Coord{X: 43, Y: 42}, foods, true},
		{"true with x - 1", models.Coord{X: 41, Y: 42}, foods, true},
		{"true with y + 1", models.Coord{X: 42, Y: 43}, foods, true},
		{"true with y - 1", models.Coord{X: 42, Y: 41}, foods, true},
		{"false with same points", models.Coord{X: 42, Y: 42}, foods, false},
		{"false with y+1 and x+1", models.Coord{X: 43, Y: 43}, foods, false},
		{"false with same x but y + 2", models.Coord{X: 42, Y: 44}, foods, false},
		{"false with same Y but x + 2", models.Coord{X: 44, Y: 42}, foods, false},
		{"false far away", models.Coord{X: 1, Y: 1}, foods, false},
		{"false food empty", models.Coord{X: 42, Y: 42}, foodsEmpty, false},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, isFoodNextTo(tc.head, tc.food))

		})
	}

}

func TestRemoveNotPossibleMoveFromObstacleNextTo(t *testing.T) {
	myHead := models.Coord{X: 42, Y: 42}
	obstacleLeft := models.Coord{X: 41, Y: 42}
	obstacleRight := models.Coord{X: 43, Y: 42}
	obstacleUp := models.Coord{X: 42, Y: 43}
	obstacleDown := models.Coord{X: 42, Y: 41}

	possibleMovesFull := map[string]bool{
		"left":  true,
		"right": true,
		"down":  true,
		"up":    true,
	}
	possibleMovesLeftFalse := map[string]bool{
		"left":  false,
		"right": true,
		"down":  true,
		"up":    true,
	}
	possibleMovesRightFalse := map[string]bool{
		"left":  true,
		"right": false,
		"down":  true,
		"up":    true,
	}
	possibleMovesDownFalse := map[string]bool{
		"left":  true,
		"right": true,
		"down":  false,
		"up":    true,
	}
	possibleMovesUpFalse := map[string]bool{
		"left":  true,
		"right": true,
		"down":  true,
		"up":    false,
	}

	tt := []struct {
		Name           string
		head           models.Coord
		obstacle       models.Coord
		expectedResult map[string]bool
	}{
		{"full with obstacle on top", myHead, obstacleUp, possibleMovesUpFalse},
		{"full with obstacle on left", myHead, obstacleLeft, possibleMovesLeftFalse},
		{"full with obstacle on bottom", myHead, obstacleDown, possibleMovesDownFalse},
		{"full with obstacle on right", myHead, obstacleRight, possibleMovesRightFalse},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, removeNotPossibleMoveFromObstacleNextTo(possibleMovesFull, tc.head, tc.obstacle))
			possibleMovesFull = map[string]bool{
				"left":  true,
				"right": true,
				"down":  true,
				"up":    true,
			}

		})
	}
}
