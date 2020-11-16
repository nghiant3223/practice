package ds

type Set struct {
	items map[*Node]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.items = make(map[*Node]struct{})
	return s
}

func (s *Set) Add(node *Node) {
	s.items[node] = struct{}{}
}

func (s *Set) Contain(node *Node) bool {
	_, ok := s.items[node]
	return ok
}

func (s *Set) Remove(node *Node) {
	delete(s.items, node)
}
