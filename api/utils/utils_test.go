package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {

	tt := []struct {
		Name           string
		number         int
		expectedResult int
	}{
		{"absolute 100", 100, 100},
		{"absolute -100", -100, 100},
		{"absolute 0", 0, 0},
		{"absolute 1", 1, 1},
		{"absolute -1", -1, 1},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, Abs(tc.number))
		})
	}
}
