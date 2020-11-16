package symetrictree

import "github.com/nghiant3223/gopractice/alg/tree/ds"

// Visit: https://leetcode.com/problems/symmetric-tree/

func isSymmetric(root *ds.Node) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(t1 *ds.Node, t2 *ds.Node) bool {
	if t1 == nil && t2 != nil {
		return false
	}
	if t1 != nil && t2 == nil {
		return false
	}
	if t1 == nil && t2 == nil {
		return true
	}
	return t1.Value == t2.Value &&
		isMirror(t1.Left, t2.Right) &&
		isMirror(t1.Right, t2.Left)
}
