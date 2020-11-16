package uniquebst2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	trees := generateTrees(3)
	assert.Equal(t, 5, len(trees))
}
