package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNextTo(t *testing.T) {
	coord := Coord{
		X: 42,
		Y: 42,
	}

	tt := []struct {
		Name           string
		OtherPoint     Coord
		expectedResult bool
	}{
		{"true with x + 1", Coord{X: 43, Y: 42}, true},
		{"true with x - 1", Coord{X: 41, Y: 42}, true},
		{"true with y + 1", Coord{X: 42, Y: 43}, true},
		{"true with y - 1", Coord{X: 42, Y: 41}, true},
		{"false with same points", Coord{X: 42, Y: 42}, false},
		{"false with y+1 and x+1", Coord{X: 43, Y: 43}, false},
		{"false with same x but y + 2", Coord{X: 42, Y: 44}, false},
		{"false with same Y but x + 2", Coord{X: 44, Y: 42}, false},
		{"false far away", Coord{X: 1, Y: 1}, false},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, coord.IsNextTo(tc.OtherPoint))
		})
	}
}
