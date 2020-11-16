package sumroot2leafnumbers

import "github.com/nghiant3223/gopractice/alg/tree/ds"

// ================== Recursive ==================

func sumNumbers(root *ds.Node) int {
	base, sum := 0, 0
	sumNumbersUtil(root, base, &sum)
	return sum
}

func sumNumbersUtil(root *ds.Node, base int, sum *int) {
	if root == nil {
		return
	}
	number := base*10 + root.Value
	if root.IsLeaf() {
		*sum += number
		return
	}
	sumNumbersUtil(root.Left, number, sum)
	sumNumbersUtil(root.Right, number, sum)
}

// ================== Iterative ==================

func sumNumbers2(root *ds.Node) int {
	if root == nil {
		return 0
	}

	sum := 0
	nodeNum := make(map[*ds.Node]int)
	visited := ds.NewSet()
	s := ds.NewStack()

	s.Push(root)
	nodeNum[root] = root.Value
	for !s.IsEmpty() {
		top := s.Top()
		if top.IsLeaf() {
			s.Pop()
			visited.Add(top)
			sum += nodeNum[top]
			continue
		}
		leftVisited := top.Left == nil || visited.Contain(top.Left)
		rightVisited := top.Right == nil || visited.Contain(top.Right)
		if leftVisited && rightVisited {
			s.Pop()
			visited.Add(top)
			continue
		}
		if top.Right != nil {
			nodeNum[top.Right] += nodeNum[top]*10 + top.Right.Value
			s.Push(top.Right)
		}
		if top.Left != nil {
			nodeNum[top.Left] += nodeNum[top]*10 + top.Left.Value
			s.Push(top.Left)
		}
	}

	return sum
}
