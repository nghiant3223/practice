package allnodeatlevel2

import (
	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

var ans [][]int

func levelOrderBottom(root *ds.Node) [][]int {
	levelOrderBottomUtil(root, 0)
	return ans
}

func levelOrderBottomUtil(root *ds.Node, level int) {
	if root == nil {
		return
	}
	ansLength := len(ans)
	if level == ansLength-1 {
		nodesIndex := ansLength - level - 1
		ans[nodesIndex] = append(ans[nodesIndex], root.Value)
	} else if level == len(ans) {
		nodes := []int{root.Value}
		ans = append([][]int{nodes}, ans...)
	}
	levelOrderBottomUtil(root.Left, level+1)
	levelOrderBottomUtil(root.Right, level+1)
}
