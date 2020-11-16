package goodnodescount

import (
	"testing"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	node3 := &ds.Node{Value: 3}
	node1 := &ds.Node{Value: 1}
	node4 := &ds.Node{Value: 4}
	node31 := &ds.Node{Value: 3}
	node11 := &ds.Node{Value: 1}
	node5 := &ds.Node{Value: 5}

	node3.Left = node1
	node3.Right = node4
	node1.Left = node31
	node4.Left = node11
	node4.Right = node5

	ans := goodNodes(node3)
	assert.Equal(t, 4, ans)
}
