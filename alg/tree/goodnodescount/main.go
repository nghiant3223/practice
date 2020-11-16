package goodnodescount

import (
	"math"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

// Visit: https://leetcode.com/problems/count-good-nodes-in-binary-tree/

func goodNodes(root *ds.Node) int {
	return traverse(root, math.MinInt32)
}

func traverse(node *ds.Node, maxSoFar int) int {
	if node == nil {
		return 0
	}
	contribution := 0
	if node.Value >= maxSoFar {
		contribution = 1
	}
	if node.Value > maxSoFar {
		maxSoFar = node.Value
	}
	return contribution + traverse(node.Left, maxSoFar) + traverse(node.Right, maxSoFar)
}
