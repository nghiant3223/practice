package main

import (
	"github.com/nghiant3223/gopractice/alg/linkedlist/ds"
	"github.com/nghiant3223/gopractice/alg/linkedlist/util"
)

// 1 2 3

func reverseLinkedList(root *ds.Node) *ds.Node {
	if root == nil {
		return nil
	}
	if root.Next == nil {
		return root
	}

	// Make 3 -> 2, 2 -> NULL
	if root.Next.Next == nil {
		newHead := root.Next
		newHead.Next = root
		root.Next = nil
		return newHead
	}

	// Reverse tail first
	newHead := reverseLinkedList(root.Next)

	// Make 2 -> 1, 1 -> NULL
	root.Next.Next = root
	root.Next = nil

	return newHead
}

func reverseLinkedListIterative(root *ds.Node) *ds.Node {
	var prev *ds.Node
	cur := root
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	root = prev
	return root
}

func main() {
	node1 := &ds.Node{Value: 1}
	node2 := &ds.Node{Value: 2}
	node3 := &ds.Node{Value: 3}
	node4 := &ds.Node{Value: 4}
	node5 := &ds.Node{Value: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	//newHead := reverseLinkedList(node1)
	//util.PrintLinkedList(newHead)
	newHead := reverseLinkedListIterative(node1)
	util.PrintLinkedList(newHead)
}
