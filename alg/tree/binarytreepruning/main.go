package binarytreepruning

import "github.com/nghiant3223/gopractice/alg/tree/ds"

// Visit: https://leetcode.com/problems/binary-tree-pruning/

const (
	left  = "l"
	right = "r"
)

func pruneTree(root *ds.Node) *ds.Node {
	if root == nil {
		return nil
	}

	leftOk := pruneNode(root, root.Left, left)
	rightOk := pruneNode(root, root.Right, right)

	if leftOk || rightOk {
		return root
	}
	if !shouldPrune(root) {
		return root
	}
	return nil
}

func pruneNode(parent *ds.Node, root *ds.Node, direction string) bool {
	if root == nil {
		return false
	}

	leftOk := pruneNode(root, root.Left, left)
	rightOk := pruneNode(root, root.Right, right)

	if leftOk || rightOk {
		return true
	}
	if !shouldPrune(root) {
		return true
	}

	if direction == left {
		parent.Left = nil
	} else {
		parent.Right = nil
	}
	return false
}

func shouldPrune(root *ds.Node) bool {
	if root.Left != nil && root.Left.Value == 1 {
		return false
	}
	if root.Right != nil && root.Right.Value == 1 {
		return false
	}
	return root.Value == 0
}
