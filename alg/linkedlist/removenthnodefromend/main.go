package removenthnodefromend

import "github.com/nghiant3223/gopractice/alg/linkedlist/ds"

func removeNthNodeFromEnd(head *ds.Node, n int) *ds.Node {
	prevLookup := make(map[int]*ds.Node)
	var prev *ds.Node
	tmp := head
	length := 0
	for tmp != nil {
		prevLookup[length] = prev
		prev = tmp
		tmp = tmp.Next
		length++
	}
	prevI := length - n
	if prevI <= 0 {
		head = head.Next
		return head
	}
	prevToNthNode := prevLookup[prevI]
	prevToNthNode.Next = prevToNthNode.Next.Next
	return head
}
