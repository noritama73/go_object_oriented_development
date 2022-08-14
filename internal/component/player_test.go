package component

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	t.Run("戦士", func(t *testing.T) {
		warrior := NewPlayer(PlayerTypeWarrior)
		cases := map[string]struct {
			expect int
			actual int
		}{
			"LEVEL": {1, warrior.Status.GetLEVEL()},
		}
		for _, c := range cases {
			assert.Equal(t, c.expect, c.actual)
		}
	})
}
