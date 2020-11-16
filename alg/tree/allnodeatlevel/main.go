package main

import (
	"fmt"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

func allNodeAtLevel(root *ds.Node, level int) []*ds.Node {
	q := &ds.Queue{}
	if level == 0 {
		return []*ds.Node{root}
	}
	q.Enqueue(root)
	currentLevel := 0
	for !q.IsEmpty() {
		currentLevel++
		var nodes []*ds.Node
		for !q.IsEmpty() {
			head := q.Dequeue()
			if head.Left != nil {
				nodes = append(nodes, head.Left)
			}
			if head.Right != nil {
				nodes = append(nodes, head.Right)
			}
		}
		if currentLevel == level {
			return nodes
		}
		for len(nodes) != 0 {
			q.Enqueue(nodes[0])
			nodes = nodes[1:]
		}
	}
	return nil
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

	for _, n := range allNodeAtLevel(node1, 2) {
		fmt.Printf("%v ", n)
	}
}
