package ds

type Stack struct {
	elems []interface{}
}

func (s *Stack) Push(e interface{}) {
	s.elems = append(s.elems, e)
}

func (s *Stack) Pop() interface{} {
	top := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Stack) Count() int {
	return len(s.elems)
}
