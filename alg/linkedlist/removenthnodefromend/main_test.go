package removenthnodefromend

import (
	"testing"

	"github.com/nghiant3223/gopractice/alg/linkedlist/ds"
	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  []int
		in2  int
		out  []int
	}{
		{
			name: "1",
			in1:  []int{0, 1, 2, 3},
			in2:  1,
			out:  []int{0, 1, 2},
		},
		{
			name: "2",
			in1:  []int{0, 1, 2, 3},
			in2:  2,
			out:  []int{0, 1, 3},
		},
		{
			name: "3",
			in1:  []int{0, 1, 2, 3},
			in2:  3,
			out:  []int{0, 2, 3},
		},
		{
			name: "4",
			in1:  []int{0, 1, 2, 3},
			in2:  4,
			out:  []int{1, 2, 3},
		},
	} {
		head := ds.ParseLinkedList(tc.in1)
		head = removeNthNodeFromEnd(head, tc.in2)
		assertEqual(t, head, tc.out)
	}
}

func assertEqual(t *testing.T, node *ds.Node, numbers []int) {
	i := 0
	for node != nil {
		assert.Equal(t, node.Value, numbers[i])
		node = node.Next
		i++
	}
	assert.Equal(t, len(numbers), i)
}
