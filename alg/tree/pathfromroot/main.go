package main

import (
	"log"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

func pathFromRootToNode(root *ds.Node, node *ds.Node, path *[]*ds.Node) bool {
	if root == nil {
		return false
	}
	// `node` found
	if root == node {
		*path = append(*path, root)
		return true
	}
	// ancestor of `node`
	if pathFromRootToNode(root.Left, node, path) || pathFromRootToNode(root.Right, node, path) {
		*path = append(*path, root)
		return true
	}
	return false
}

func main() {
	node1 := &ds.Node{Value: 1}
	node2 := &ds.Node{Value: 2}
	node3 := &ds.Node{Value: 3}
	node4 := &ds.Node{Value: 4}
	node5 := &ds.Node{Value: 5}
	node6 := &ds.Node{Value: 6}
	node7 := &ds.Node{Value: 7}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	var path []*ds.Node
	if !pathFromRootToNode(node1, node7, &path) {
		log.Fatal("not found")
		return
	}
	for _, n := range path {
		log.Print(n)
	}
}
