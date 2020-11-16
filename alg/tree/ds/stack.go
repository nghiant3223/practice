package ds

type Stack struct {
	items []*Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(node *Node) {
	s.items = append(s.items, node)
}

func (s *Stack) Pop() *Node {
	topItemIndex := len(s.items) - 1
	top := s.items[topItemIndex]
	s.items = s.items[0:topItemIndex]
	return top
}

func (s *Stack) Top() *Node {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
