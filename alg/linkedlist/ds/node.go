package ds

type Node struct {
	Value int
	Next  *Node
}

func NewNode(val int) *Node {
	return &Node{Value: val}
}
