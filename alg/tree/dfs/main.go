package main

import (
	"log"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

var visited map[*ds.Node]struct{}

func depthFirstSearch(root *ds.Node) {
	s := ds.NewStack()
	s.Push(root)
	visited = make(map[*ds.Node]struct{})
	for !s.IsEmpty() {
		top := s.Pop()
		visit(top)
		_, ok := visited[top.Right]
		if top.Right != nil && !ok {
			s.Push(top.Right)
		}
		_, ok = visited[top.Left]
		if top.Left != nil && !ok {
			s.Push(top.Left)
		}
	}
}

func visit(node *ds.Node) {
	visited[node] = struct{}{}
	log.Println(node.Value)
}

func main() {
	root := ds.ParseTree([]int{1, 2, 3, 4, 5, 6, 7, ds.Nil, 8, ds.Nil, ds.Nil, 9, 10, ds.Nil, ds.Nil})
	depthFirstSearch(root)
}
