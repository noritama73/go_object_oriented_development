package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		assert.Equal(t, 1 + 2, Add(1, 2))
	})
}