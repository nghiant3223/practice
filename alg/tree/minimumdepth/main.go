package minimumdepth

import "github.com/nghiant3223/gopractice/alg/tree/ds"

func minimumDepth(root *ds.Node) int {
	if root == nil {
		return 0
	}
	return min(minimumDepth(root.Left), minimumDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
