package ds

type Queue struct {
	items []*Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(n *Node) {
	q.items = append(q.items, n)
}

func (q *Queue) Dequeue() *Node {
	head := q.items[0]
	q.items = q.items[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
