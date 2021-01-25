package checkpalindrome

import (
	"testing"

	"github.com/nghiant3223/gopractice/alg/linkedlist/ds"
	"github.com/stretchr/testify/assert"
)

func TestHappy(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []int
		out  bool
	}{
		{
			name: "1",
			in:   []int{1, 2, 3, 4, 3, 2, 1},
			out:  true,
		},
		{
			name: "2",
			in:   []int{1, 2, 3, 3, 2, 1},
			out:  true,
		},
		{
			name: "3",
			in:   []int{1, 2, 4, 4, 3, 2, 1},
			out:  false,
		},
	} {
		head := ds.ParseLinkedList(tc.in)
		out := isPalindrome(head)
		assert.Equal(t, tc.out, out)
	}
}
