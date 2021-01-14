package ds

const Nil = -1

func ParseTree(nums []int) *Node {
	length := len(nums)
	nodes := make([]*Node, length)
	for i, num := range nums {
		var node *Node
		if num != Nil {
			node = NewNode(num)
		}
		nodes[i] = node
	}
	for i := 0; 2*i+2 < length; i++ {
		node := nodes[i]
		node.Left = nodes[2*i+1]
		node.Right = nodes[2*i+2]
	}
	root := nodes[0]
	return root
}
