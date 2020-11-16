package allnodeatlevel2

import (
	"testing"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	node3 := &ds.Node{Value: 3}
	node9 := &ds.Node{Value: 9}
	node20 := &ds.Node{Value: 20}
	node15 := &ds.Node{Value: 15}
	node7 := &ds.Node{Value: 7}

	node3.Left = node9
	node3.Right = node20
	node20.Left = node15
	node20.Right = node7

	ans := levelOrderBottom(node3)
	assert.Equal(t, [][]int{{15, 7}, {9, 20}, {3}}, ans)
}
