package ds

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func (n *Node) Equal(m *Node) bool {
	return n.Value == m.Value
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}