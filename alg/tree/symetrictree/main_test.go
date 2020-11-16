package symetrictree

import (
	"testing"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	node1 := &ds.Node{Value: 1}
	node21 := &ds.Node{Value: 2}
	node22 := &ds.Node{Value: 2}
	node31 := &ds.Node{Value: 3}
	node32 := &ds.Node{Value: 3}
	node41 := &ds.Node{Value: 4}
	node42 := &ds.Node{Value: 4}

	node1.Left = node21
	node1.Right = node22
	node21.Left = node31
	node21.Right = node41
	node22.Left = node42
	node22.Right = node32

	ans := isSymmetric(node1)
	assert.True(t, ans)
}
