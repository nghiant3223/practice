package maxsumfromroot2leaf

import "github.com/nghiant3223/gopractice/alg/tree/ds"

func maxSumFromRoot2LeafPath(root *ds.Node) (int, []int) {
	if root == nil {
		return 0, nil
	}
	sumLeft, pathLeft := maxSumFromRoot2LeafPath(root.Left)
	sumRight, pathRight := maxSumFromRoot2LeafPath(root.Right)
	if sumLeft > sumRight {
		return sumLeft + root.Value, append([]int{root.Value}, pathLeft...)
	}
	return sumRight + root.Value, append([]int{root.Value}, pathRight...)
}
