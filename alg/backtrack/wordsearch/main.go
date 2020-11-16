package main

import "fmt"

func find(word string, root *Node, i int) bool {
	if root == nil {
		return false
	}
	if i > len(word) {
		return false
	}
	if i == len(word)-1 {
		return rune(word[i]) == root.Value
	}
	if rune(word[i]) != root.Value {
		return false
	}
	leftFound := find(word, root.Left, i+1)
	if leftFound {
		return true
	}
	rightFound := find(word, root.Right, i+1)
	return rightFound
}

func main() {
	A := &Node{Value: 'A'}
	N := &Node{Value: 'N'}
	I := &Node{Value: 'I'}
	T := &Node{Value: 'T'}
	D := &Node{Value: 'D'}
	M := &Node{Value: 'M'}
	R := &Node{Value: 'R'}
	A.Left = N
	A.Right = I
	N.Left = T
	N.Right = D
	I.Left = M
	I.Right = R

	root := A
	word := "AND"
	found := find(word, root, 0)
	fmt.Println(found)
}
