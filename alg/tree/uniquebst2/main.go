package uniquebst2

import (
	"fmt"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
)

func generateTrees(n int) []*ds.Node {
	var roots []*ds.Node
	shapeSet := make(map[string]struct{})
	perms := generatePermutations(n)
	for _, perm := range perms {
		root := constructBinaryTree(perm)
		shape := getShape(root)
		if _, ok := shapeSet[shape]; ok {
			continue
		}
		shapeSet[shape] = struct{}{}
		roots = append(roots, root)
	}
	return roots
}

func constructBinaryTree(perms []int) *ds.Node {
	root := &ds.Node{Value: perms[0]}
	for i := 1; i < len(perms); i++ {
		value := perms[i]
		newNode := &ds.Node{Value: value}
		insertNodeToBST(root, newNode)
	}
	return root
}

func insertNodeToBST(node *ds.Node, newNode *ds.Node) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNodeToBST(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNodeToBST(node.Right, newNode)
		}
	}
}

func getShape(node *ds.Node) string {
	q := ds.NewQueue()
	q.Enqueue(node)
	s := ""
	for !q.IsEmpty() {
		head := q.Dequeue()
		s += fmt.Sprintf("%d,", head.Value)
		if head.Left == nil {
			s += fmt.Sprintf("NULL,")
		} else {
			q.Enqueue(head.Left)
		}
		if head.Right == nil {
			s += fmt.Sprintf("NULL,")
		} else {
			q.Enqueue(head.Right)
		}
	}
	return s
}

func generatePermutations(n int) [][]int {
	var ans [][]int
	var curs []int
	generatePermutationsUtil(n, &ans, curs)
	return ans
}

func generatePermutationsUtil(n int, ans *[][]int, curs []int) {
	if len(curs) == n {
		*ans = append(*ans, curs)
		return
	}
	for i := 1; i <= n; i++ {
		if contains(curs, i) {
			continue
		}
		generatePermutationsUtil(n, ans, append(curs, i))
	}
}

func contains(s []int, t int) bool {
	for _, v := range s {
		if v == t {
			return true
		}
	}
	return false
}
