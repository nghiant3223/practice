package isvalidbst

import (
	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

func isValidBST(root *ds.Node) bool {
	if root == nil {
		return true
	}
	leftOk := isValidBST(root.Left)
	if !leftOk {
		return false
	}
	ok := visitNode(root)
	if !ok {
		return false
	}
	rightOk := isValidBST(root.Right)
	return rightOk
}

var prev *ds.Node

func visitNode(node *ds.Node) bool {
	defer func() { prev = node }()
	if prev == nil {
		return true
	}
	return prev.Value < node.Value
}
