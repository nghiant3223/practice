package maxancestordiff

import (
	"math"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

func maxAncestorDiff(root *ds.Node) int {
	_, _, maxDiff := maxAncestorDiffHelper(root, math.MaxInt32, math.MinInt32)
	return maxDiff
}

func maxAncestorDiffHelper(root *ds.Node, parentMin int, parentMax int) (int, int, int) {
	if root == nil {
		return math.MaxInt32, math.MinInt32, 0
	}
	currentMin := min(parentMin, root.Value)
	currentMax := max(parentMax, root.Value)
	minLeft, maxLeft, _ := maxAncestorDiffHelper(root.Left, currentMin, currentMax)
	minRight, maxRight, _ := maxAncestorDiffHelper(root.Right, currentMin, currentMax)
	successorMin := min(root.Value, min(minLeft, minRight))
	successorMax := max(root.Value, min(maxLeft, maxRight))
	maxDiffAtCurrentNode := max(abs(successorMin-currentMax), abs(successorMax-currentMin))
	return successorMin, successorMax, maxDiffAtCurrentNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
