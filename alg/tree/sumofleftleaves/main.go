package sumofleftleaves

import "github.com/nghiant3223/gopractice/alg/tree/ds"

func sumOfLeftLeaves(root *ds.Node) int {
	return dfs(root, false)
}

func dfs(root *ds.Node, isLeft bool) int {
	if root == nil {
		return 0
	}
	if root.IsLeaf() && isLeft {
		return root.Value
	}
	return dfs(root.Left, true) + dfs(root.Right, false)
}
