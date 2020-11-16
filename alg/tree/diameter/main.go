package main

import (
	"fmt"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

// Given a binary tree, find the longest path between two any node in the tree

// Idea: diameter of a tree is the sum depth of two deepest nodes
// For left subtree and right subtree, calculate the depth of deepest node and its best diameter
// To get the best diameter overall, we get max of:
// 1. left subtree diameter,
// 2. right subtree diameter,
// 3. sum depth of deepest node in the left and that 	in the right subtree

func dfs(root *ds.Node, depth int) (maxDepth int, diameter int) {
	var leftMaxDepth, leftDiameter, rightMaxDepth, rightDiameter int
	if root.Left != nil {
		leftMaxDepth, leftDiameter = dfs(root.Left, depth + 1)
	}
	if root.Right != nil {
		rightMaxDepth, rightDiameter = dfs(root.Right, depth + 1)
	}
	return max(leftMaxDepth, rightMaxDepth) + 1, max3(leftDiameter, rightDiameter, leftMaxDepth+rightMaxDepth)
}

func max3(a, b, c int) int {
	return max(a, max(b, c))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n1 := &ds.Node{Value: 1}
	n2 := &ds.Node{Value: 2}
	n3 := &ds.Node{Value: 3}
	n4 := &ds.Node{Value: 4}
	n5 := &ds.Node{Value: 5}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	_, ans := dfs(n1, 0)
	fmt.Println(ans)
}
